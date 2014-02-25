// Copyright (c) 2014 Bertrand Janin <b@janin.com>
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"regexp"
	"strings"
)

var (
	reStop = regexp.MustCompile(`^st[aho]+p`)
	reShhh = regexp.MustCompile(`^s+[sh]+`)
)

type ShutUpModule struct{}

func (module ShutUpModule) PrivMsg(msg *PrivMsg) {}

func isShutUpRequest(msg *PrivMsg) bool {
	body := strings.ToLower(msg.Body)
	println(body)
	if reStop.MatchString(body) {
		return true
	}
	if reShhh.MatchString(body) {
		return true
	}
	if strings.HasPrefix(body, "shut up") {
		return true
	}
	return false
}

func ShutUpCommand(where string, params []string) {
	SendToMinion("shutup")
	privMsg(where, "ok...")
}

func (module ShutUpModule) Init() {
	RegisterCommand(Command{
		Name:           "shutup",
		ToggleFunction: isShutUpRequest,
		Function:       ShutUpCommand,
		Addressed:      true,
		AllowDirect:    false,
		AllowChannel:   true,
	})
}
