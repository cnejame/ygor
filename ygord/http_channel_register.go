// Copyright 2014-2015, Truveris Inc. All Rights Reserved.
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type ChannelRegisterHandler struct {
	*Server
}

type ChannelRegisterRequest struct {
	ChannelID string
}

type ChannelRegisterResponse struct {
	ClientID string
}

func (handler *ChannelRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, err := auth(r)
	if err != nil {
		errorHandler(w, "Authentication failed", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	input := &ChannelRegisterRequest{}
	err = decoder.Decode(input)
	if err != nil {
		errorHandler(w, "Failed to decode input JSON", err)
		return
	}

	clientID := handler.Server.RegisterClient(username, "#"+input.ChannelID)

	w.Header().Set("Content-Type", "application/json")

	select {
	case <-time.After(time.Second * 2):
	}

	JSONHandler(w, ChannelRegisterResponse{ClientID: clientID})
}
