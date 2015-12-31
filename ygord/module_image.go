// Copyright 2015, Truveris Inc. All Rights Reserved.
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

// ImageModule controls the 'image' command.
type ImageModule struct {
	*Server
}

// PrivMsg is the message handler for user 'image' requests.
func (module *ImageModule) PrivMsg(srv *Server, msg *InputMessage) {
	usage := "usage: image url [end]"

	// Validate the command's usage, and get back a map representing the media
	// item that was passed, along with it's start and end bounds.
	mediaItem, err := parseArgList(msg.Args)
	if err != nil {
		srv.Reply(msg, usage)
		return
	}

	media, err := NewMedia(srv, mediaItem, "imageTrack", true, true,
		[]string{
			"vimeo",
			"youtube",
			"video",
			"img",
			"web",
		})
	if err != nil {
		srv.Reply(msg, err.Error())
		return
	}

	// Send the command to the connected minions.
	srv.SendToChannelMinions(msg.ReplyTo, ClientCommand{"image", media})
}

// Init registers all the commands for this module.
func (module ImageModule) Init(srv *Server) {
	srv.RegisterCommand(Command{
		Name:            "image",
		PrivMsgFunction: module.PrivMsg,
		Addressed:       true,
		AllowPrivate:    false,
		AllowChannel:    true,
	})
}
