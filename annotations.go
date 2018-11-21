package main

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
