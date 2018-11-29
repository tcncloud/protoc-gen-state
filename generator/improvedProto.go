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
  field                 *gp.FieldDescriptorProto
  parentMessagesString  string // all the message names leading up to this message name. Empty most of the time
  packageName           string
  file                  *gp.FileDescriptorProto
  message               *ImprovedMessageDescriptor
}


type ImprovedMessageDescriptor struct {
  message         *gp.DescriptorProto
  fields          []*ImprovedFieldDescriptor
  parentMessage   *ImprovedMessageDescriptor
  childMessages   []*ImprovedMessageDescriptor //getNestedMessages
  packageName     string
  file            *gp.FileDescriptorProto
}


func FieldDescriptorToImproved(field *gp.FieldDescriptorProto, files []*gp.FileDescriptorProto) *ImprovedFieldDescriptor {
  var parentMsg *gp.DescriptorProto
  var foundFile *gp.FileDescriptorProto

  for _, file := range files {
    packageName := file.GetPackage()
    for _, message := range file.GetMessageType() {
      msgName := fmt.Sprintf(".%s.%s", packageName, message.GetName())

      if msgName == field.GetTypeName() {
        // found the parent message
        parentMsg = message
        foundFile = file
        break;
      }

      // check nested types too
      nested := message.GetNestedType()
      win, desc := checkNestedType(msgName, nested, field.GetTypeName())
      if win {
        // found the parent message
        parentMsg = desc
        foundFile = file
        break;
      }
    }
  }

  parentMessagesString := "" // all the message names leading up to this message name. Empty most of the time


  return &ImprovedFieldDescriptor{
    field: field,
    packageName: foundFile.GetPackage(),
    file: foundFile,
    parentMessagesString: parentMessagesString,
    message: MessageDescriptorToImproved(parentMsg),
  }
}

func MessageDescriptorToImproved(message *gp.DescriptorProto) *ImprovedMessageDescriptor {
  return &ImprovedMessageDescriptor{
    message : message,
    fields : nil,
    parentMessage: nil,
    childMessages: nil,
    packageName: "",
    file: nil,
  }
}
