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
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	// https://godoc.org/github.com/golang/protobuf/protoc-gen-go/plugin#CodeGeneratorRequest
	var req plugin_go.CodeGeneratorRequest

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(fmt.Errorf("got error reading from stdin: %v", err))
	}
	if err := proto.Unmarshal(data, &req); err != nil {
		panic(fmt.Errorf("got error unmarshaling request: %v", err))
	}
	files, err, internalErr := generate(req.GetFileToGenerate(), req.GetProtoFile())
	if internalErr != nil {
		panic(fmt.Errorf("error generating: %v", err))
	}
	resp := plugin_go.CodeGeneratorResponse{
		Error: func() (out *string) {
			if err != nil {
				*out = fmt.Sprintf("%s", err)
				return
			}
			return
		}(),
		File: func() (out []*plugin_go.CodeGeneratorResponse_File) {
			if err != nil {
				return
			}
			for _, f := range files {
				out = append(out, &plugin_go.CodeGeneratorResponse_File{
					Name:    &f.Name,
					Content: &f.Content,
				})
			}
			return
		}(),
	}
	data, err = proto.Marshal(&resp)
	if err != nil {
		panic(fmt.Errorf("error marshaling the response: %v", err))
	}
	if _, err := os.Stdout.Write(data); err != nil {
		panic(fmt.Errorf("error writing to std out: %v", err))
	}
}
