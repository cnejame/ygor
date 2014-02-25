// Copyright (c) 2014 Bertrand Janin <b@janin.com>
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"strings"
)

type XombreroModule struct{}

func (module XombreroModule) PrivMsg(msg *PrivMsg) {}

func XombreroFunc(msg *PrivMsg) {
	if len(msg.Args) == 0 {
		privMsg(msg.ReplyTo, "usage: xombrero [command [param ...]]")
		return
	}

	SendToMinion("xombrero " + strings.Join(msg.Args, " "))
	privMsg(msg.ReplyTo, "sure")
}

func (module XombreroModule) Init() {
	RegisterCommand(Command{
		Name:         "xombrero",
		Function:     XombreroFunc,
		Addressed:    true,
		AllowDirect:  false,
		AllowChannel: true,
	})
}
