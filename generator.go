// Copyright 2017, TCN Inc.
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of TCN Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"errors"
	"fmt"
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type File struct {
	Name    string
	Content string
}

func generate(filepaths []string, protos []*gp.FileDescriptorProto) ([]*File, error) {
	// the two messages we'll be reading
	var stateMessage *gp.DescriptorProto
	var customMessage *gp.DescriptorProto

	// throw an error if len(filepaths) > 1
	if len(filepaths) > 1 {
		return nil, errors.New("Multiple file inputs detected. This plugin is designed to generate redux state from a single proto file")
	}

	// find the file descriptor and package name for the state message
	// var statePackageName string
	var stateFile *gp.FileDescriptorProto
	for _, p := range protos {
		if p.GetName() == filepaths[0] {
			// statePackageName = p.GetPackage()
			stateFile = p
		}
	}

	messageCount := len(stateFile.GetMessageType())
	// at least one message must be defined or we can't generate anything
	if messageCount == 0 {
		return nil, errors.New("No messages defined in state proto: " + stateFile.GetName() + ". Please include a ReduxState or CustomActions message.")
	}
	// there are only 3 message allowed in the state message
	if messageCount > 3 {
		return nil, errors.New("Too many messages defined in state proto: " + stateFile.GetName() + ". Only ReduxState, CustomActions, and ExternalLink messages allowed.")
	}

	// enforce that the messages provided are the allowed messages
	// TODO look into removing ExternalLink message
	allowedNames := []string{"ReduxState", "CustomActions", "ExternalLink"}
	for _, m := range stateFile.GetMessageType() {
		if !contains(allowedNames, m.GetName()) {
			return nil, errors.New("Bad message name encountered: " + m.GetName() + ". Only ReduxState, CustomActions, and ExternalLink messages allowed in state message.")
		} else if m.GetName() == "ReduxState" {
			stateMessage = m
		} else if m.GetName() == "CustomActions" {
			customMessage = m
		}
	}

	// gather the configuration annotations, throw if they're required
	debounce, err := GetFileExtensionInt(stateFile, "debounce")
	defaultTimeout, err := GetFileExtensionInt(stateFile, "default_timeout")
	defaultRetries, err := GetFileExtensionInt(stateFile, "default_retries")
	port, err := GetFileExtensionInt(stateFile, "port")
	debug, err := GetFileExtensionBool(stateFile, "debug")
	protocTsPath, err := GetFileExtensionString(stateFile, "protoc_ts_path")
	hostname, err := GetFileExtensionString(stateFile, "hostname")
	hostnameLocation, err := GetFileExtensionString(stateFile, "hostname_location")
	authTokenLocation, err := GetFileExtensionString(stateFile, "auth_token_location")

	if err != nil {
		return nil, fmt.Errorf("Error encountered while parsing file level annotations: %v", err)
	}

	stateFields := []*gp.FieldDescriptorProto{}
	customFields := []*gp.FieldDescriptorProto{}
	messageFiles := []*gp.FileDescriptorProto{}
	serviceFiles := []*gp.FileDescriptorProto{}

	// populate messageFiles and serviceFiles
	for _, proto := range protos {
		for _, d := range proto.Dependency {
			if proto.GetName() == d {
				continue
			} else {
				if len(proto.GetMessageType()) > 0 && !containsFile(messageFiles, proto) {
					messageFiles = append(messageFiles, proto)
				}
				if len(proto.GetService()) > 0 && !containsFile(serviceFiles, proto) {
					serviceFiles = append(serviceFiles, proto)
				}
			}
		}
	}

	// populate the stateFields by looking at the ReduxState message
	for _, field := range stateMessage.GetField() {
		stateFields = append(stateFields, field)
	}
	// populate the customFields by looking at the CustomActions message
	for _, field := range customMessage.GetField() {
		customFields = append(customFields, field)
	}

	// list of output files
	out := make([]*File, 0)

	// create aggregate for each unique package name
	out = append(out, CreateAggregateByPackage(messageFiles, protocTsPath, stateFile.GetPackage())...)

	// create state file
	statePb, err := CreateStateFile(stateFields)
	if err != nil {
		return nil, fmt.Errorf("Error generating state_pb file: %v", err)
	}
	out = append(out, statePb)

	// create action file
	actionPb, err := CreateActionFile(stateFields, customFields, serviceFiles)
	if err != nil {
		return nil, fmt.Errorf("Error generating actions_pb file: %v", err)
	}
	out = append(out, actionPb)

	// create reducer file
	reducerPb, err := CreateReducerFile(stateFields)
	if err != nil {
		return nil, fmt.Errorf("Error generating reducer_pb file: %v", err)
	}
	out = append(out, reducerPb)

	// create epic file
	epicPb, err := CreateEpicFile(stateFields, customFields, serviceFiles, defaultTimeout, defaultRetries, authTokenLocation, hostnameLocation, hostname, port, debounce, debug)
	if err != nil {
		return nil, fmt.Errorf("Error generating actions_pb file: %v", err)
	}
	out = append(out, epicPb)

	// create toMessage file
	toMessagePb, err := CreateToMessageFile(serviceFiles, protos, protocTsPath)
	if err != nil {
		return nil, fmt.Errorf("Error generating to_message_pb file: %v", err)
	}
	out = append(out, toMessagePb)

	// create message_aggregate file from the messageFiles
	typesPb, err := CreateAggregateTypesFile(messageFiles, stateFile.GetPackage())
	if err != nil {
		return nil, fmt.Errorf("Error generating protoc_types_pb file: %v", err)
	}
	out = append(out, typesPb)

	// create service_aggregate file from the serviceFiles
	servicesPb, err := CreateAggregateServicesFile(serviceFiles, protocTsPath, stateFile.GetPackage())
	if err != nil {
		return nil, fmt.Errorf("Error generating protoc_services_pb file: %v", err)
	}
	out = append(out, servicesPb)

	return out, nil
}
