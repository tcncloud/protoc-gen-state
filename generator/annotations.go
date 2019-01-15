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
	"errors"
	"github.com/golang/protobuf/proto"
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/tcncloud/protoc-gen-state/state"
)

func GetFileExtensions(desc *gp.FileDescriptorProto) (*state.StateFileOptions, error) {
  var extName *proto.ExtensionDesc
  extName = state.E_StateFileOptions

  if desc == nil || desc.GetOptions() == nil {
    return nil, errors.New("Could not get file extensions. Descriptor was empty")
  }

  if proto.HasExtension(desc.GetOptions(), extName) {
    pkg, err := proto.GetExtension(desc.GetOptions(), extName)
    if err != nil {
      return nil, err
    }
    return pkg.(*state.StateFileOptions), nil
  }
  return nil, errors.New("Could not get file extensions")
}

func GetFieldOptions(desc *gp.FieldDescriptorProto) (*state.StateFieldOptions, error) {
  var extName *proto.ExtensionDesc
  extName = state.E_StateFieldOptions

  if desc == nil || desc.GetOptions() == nil {
    return nil, errors.New("Could not get field extensions. Descriptor was empty")
  }

  if proto.HasExtension(desc.GetOptions(), extName) {
    pkg, err := proto.GetExtension(desc.GetOptions(), extName)
    if err != nil {
      return nil, err
    }
    return pkg.(*state.StateFieldOptions), nil
  }
  return nil, nil
}

func GetAnnotation(meth state.StringFieldOptions, crud Crud, repeated bool) string {
	switch crud {
	case CREATE:
		return meth.GetCreate()
	case GET:
		{
			if repeated {
				return meth.GetList()
			} else {
				return meth.GetGet()
			}
		}
	case UPDATE:
		return meth.GetUpdate()
	case DELETE:
		return meth.GetDelete()
	default:
		return ""
	}
}
