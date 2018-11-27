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

func GetFileExtensionBool(desc *gp.FileDescriptorProto, name string) (bool, error) {
	var extName *proto.ExtensionDesc
	switch name {
	case "debug":
		extName = state.E_Debug
	default:
		return false, errors.New("Unrecognized extension name: " + name + ". Maybe it's a string or int?")
	}
	return GetExtensionBool(desc, extName)
}

func GetFileExtensionInt(desc *gp.FileDescriptorProto, name string) (int64, error) {
	var extName *proto.ExtensionDesc
	switch name {
	case "default_timeout":
		extName = state.E_DefaultTimeout
	case "default_retries":
		extName = state.E_DefaultRetries
	case "debug": // its a bool though
		extName = state.E_Debug
	case "port":
		extName = state.E_Port
	case "debounce":
		extName = state.E_Debounce
	default:
		return 0, errors.New("Unrecognized extension name: " + name + ". Maybe it's a string or bool?")
	}

	return GetExtensionInt(desc, extName)
}

func GetFileExtensionString(desc *gp.FileDescriptorProto, name string) (string, error) {
	var extName *proto.ExtensionDesc
	switch name {
	case "protoc_ts_path":
		extName = state.E_ProtocTsPath
	case "hostname":
		extName = state.E_Hostname
	case "hostname_location":
		extName = state.E_HostnameLocation
	case "auth_token_location":
		extName = state.E_AuthTokenLocation
	default:
		return "", errors.New("Unrecognized extension name: " + name + ". Maybe it's an int or bool?")
	}

	return GetExtensionString(desc, extName)
}

//   ಠ_ಠ   okay
func GetExtensionInt(desc *gp.FileDescriptorProto, extName *proto.ExtensionDesc) (int64, error) {
	if desc == nil || desc.GetOptions() == nil {
		return -1, nil
	}
	if proto.HasExtension(desc.GetOptions(), extName) {
		pkg, err := proto.GetExtension(desc.GetOptions(), extName)
		if err != nil {
			return -1, errors.New("Failed to get debounce extension")
		}
		return *pkg.(*int64), nil
	}
	return -1, nil
}

//   ಠ_ಠ   okay
func GetExtensionString(desc *gp.FileDescriptorProto, extName *proto.ExtensionDesc) (string, error) {
	if desc == nil || desc.GetOptions() == nil {
		return "", nil
	}
	if proto.HasExtension(desc.GetOptions(), extName) {
		pkg, err := proto.GetExtension(desc.GetOptions(), extName)
		if err != nil {
			return "", errors.New("Failed to get debounce extension")
		}
		return *pkg.(*string), nil
	}
	return "", nil
}

//   ಠ_ಠ   okay
func GetExtensionBool(desc *gp.FileDescriptorProto, extName *proto.ExtensionDesc) (bool, error) {
	if desc == nil || desc.GetOptions() == nil {
		return false, nil
	}
	if proto.HasExtension(desc.GetOptions(), extName) {
		pkg, err := proto.GetExtension(desc.GetOptions(), extName)
		if err != nil {
			return false, errors.New("Failed to get debounce extension")
		}
		return *pkg.(*bool), nil
		// return *pkg.(*string), nil
	}
	return false, nil
}

func GetFieldOptionsString(desc *gp.FieldDescriptorProto, extName *proto.ExtensionDesc) (state.StringFieldOptions, error) {
	junk := state.StringFieldOptions{}
	if desc == nil || desc.GetOptions() == nil {
		return junk, nil
	}
	if proto.HasExtension(desc.GetOptions(), extName) {
		pkg, err := proto.GetExtension(desc.GetOptions(), extName)
		if err != nil {
			return junk, errors.New("Failed to get debounce extension")
		}
		return *pkg.(*state.StringFieldOptions), nil
	}
	return junk, nil
}

func GetFieldOptionsInt(desc *gp.FieldDescriptorProto, extName *proto.ExtensionDesc) (state.IntFieldOptions, error) {
	junk := state.IntFieldOptions{}
	if desc == nil || desc.GetOptions() == nil {
		return junk, nil
	}
	if proto.HasExtension(desc.GetOptions(), extName) {
		pkg, err := proto.GetExtension(desc.GetOptions(), extName)
		if err != nil {
			return junk, errors.New("Failed to get debounce extension")
		}
		return *pkg.(*state.IntFieldOptions), nil
	}
	return junk, nil
}

func GetFieldAnnotationInt(desc *gp.FieldDescriptorProto, extName *proto.ExtensionDesc) (int64, error) {
	if desc == nil || desc.GetOptions() == nil {
		return -1, nil
	}
	if proto.HasExtension(desc.GetOptions(), extName) {
		pkg, err := proto.GetExtension(desc.GetOptions(), extName)
		if err != nil {
			return -1, errors.New("Failed to get debounce extension")
		}
		return *pkg.(*int64), nil
	}
	return -1, nil
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
