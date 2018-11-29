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

package generator

import (
	"fmt"

	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type ImprovedFieldDescriptor struct {
	field                *gp.FieldDescriptorProto
	parentMessagesString string // all the message names leading up to this message name. Empty most of the time
	packageName          string
	file                 *gp.FileDescriptorProto
	message              *ImprovedMessageDescriptor
}

type ImprovedMessageDescriptor struct {
	message       *gp.DescriptorProto
	fields        []*ImprovedFieldDescriptor
	parentMessage *ImprovedMessageDescriptor
	packageName   string
	file          *gp.FileDescriptorProto
}

func FieldDescriptorToImproved(field *gp.FieldDescriptorProto, files []*gp.FileDescriptorProto) *ImprovedFieldDescriptor {
	var parentMsg *gp.DescriptorProto
	var foundFile *gp.FileDescriptorProto
	parentMessagesString := ""

	for _, file := range files {
		packageName := file.GetPackage()
		for _, message := range file.GetMessageType() {
			msgName := fmt.Sprintf(".%s.%s", packageName, message.GetName())

			if msgName == field.GetTypeName() {
				// found the parent message
				parentMsg = message
				foundFile = file
				break
			}

			// check nested types too
			nested := message.GetNestedType()
			win, desc, depth := checkNestedType(msgName, nested, field.GetTypeName())
			if win {
				// found the parent message
				parentMsg = desc
				foundFile = file
				parentMessagesString = depth
				break
			}
		}
	}

	return &ImprovedFieldDescriptor{
		field:                field,
		packageName:          foundFile.GetPackage(),
		file:                 foundFile,
		parentMessagesString: parentMessagesString,
		message:              MessageDescriptorToImproved(parentMsg, files),
	}
}

func MessageDescriptorToImproved(message_in *gp.DescriptorProto, files []*gp.FileDescriptorProto) *ImprovedMessageDescriptor {
	// fields := []*ImprovedFieldDescriptor{}
	// for i := 0; i < len(message_in.GetField()); i++ {
	// 	fields = append(fields, FieldDescriptorToImproved(message_in.GetField()[i], files))
	// }

	// var foundFile *gp.FileDescriptorProto
	// var foundParentMessage *gp.DescriptorProto

	// for _, file := range files {
	// 	for _, currMessage := range file.GetMessageType() {
	// 		if currMessage.GetName() == message_in.GetName() {
	// 			foundParentMessage = nil
	// 			foundFile = file
	// 			break
	// 		}

	// 		// check nested messages for our message
	// 		found, parent := FindParentMessage(currMessage, message_in)
	// 		if found {
	// 			// found the parent message
	// 			foundParentMessage = parent
	// 			foundFile = file
	// 			break
	// 		}
	// 	}
	// }

	// return &ImprovedMessageDescriptor{
	// 	message:       message_in,
	// 	fields:        fields,
	// 	parentMessage: MessageDescriptorToImproved(foundParentMessage, files),
	// 	packageName:   foundFile.GetPackage(),
	// 	file:          foundFile,
	// }
	return &ImprovedMessageDescriptor{
		message:       message_in,
		fields:        nil,
		parentMessage: nil,
		packageName:   "package.name",
		file:          nil,
	}
}

func CreateImprovedDescriptors(protos []*gp.FileDescriptorProto) []*ImprovedMessageDescriptor {
	output := []*ImprovedMessageDescriptor{}
	for _, file := range protos { // loop through files
		for _, message := range file.GetMessageType() { // loop through messages
			output = createNestedImprovedDescriptor(message, nil, file, output, protos)
		}
	}
	return output
}

func createNestedImprovedDescriptor(message *gp.DescriptorProto, prev *ImprovedMessageDescriptor, file *gp.FileDescriptorProto, fullList []*ImprovedMessageDescriptor, protos []*gp.FileDescriptorProto) []*ImprovedMessageDescriptor {
	// make our descriptor and add it to the array
	fields := []*ImprovedFieldDescriptor{} // field descriptors
	for _, f := range message.GetField() {
		fields = append(fields, FieldDescriptorToImproved(f, protos))
	}

	improvedMessage := &ImprovedMessageDescriptor{
		message:       message,
		fields:        fields,
		parentMessage: prev,
		packageName:   file.GetPackage(),
		file:          file,
	}

	fullList = append(fullList, improvedMessage)

	// recursively call this function
	if len(message.GetNestedType()) != 0 { // if we have nested types
		for _, nested := range message.GetNestedType() { // recursively get those types
			fullList = createNestedImprovedDescriptor(nested, improvedMessage, file, fullList, protos)
		}
	}

	return fullList
}

func FindImprovedFromDescriptor(improved []*ImprovedMessageDescriptor, desc *gp.DescriptorProto) *ImprovedMessageDescriptor {
	for _, i := range improved {
		if i.message == desc {
			return i
		}
	}
	return nil
}

func FindImprovedPathName(start *ImprovedMessageDescriptor) string {
	result := ""
	current := start
	for current.parentMessage != nil {
		result += "." + current.parentMessage.message.GetName()
		current = current.parentMessage
	}
	return result
}
