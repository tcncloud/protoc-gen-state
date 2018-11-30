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
	"regexp"

	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"strings"
)

type SideEffect int
type Crud int
type CludgeEffect int

const (
	REQUEST         SideEffect = 0
	SUCCESS         SideEffect = 1
	FAILURE         SideEffect = 2
	CANCEL          SideEffect = 3
	SIDE_EFFECT_MAX SideEffect = 4
)

const (
	CREATE   Crud = 0
	GET      Crud = 1
	UPDATE   Crud = 2
	DELETE   Crud = 3
	CRUD_MAX Crud = 4
	CUSTOM   Crud = 5
)

func CreatePackageAndTypeString(in string) string {
	// remove the first character if it's a period
	if in[0] == '.' {
		in = in[1:]
	}

	period := regexp.MustCompile("\\.")
	numPeriods := len(period.FindAllStringIndex(in, -1))

	// if there is only one period, the package name has no periods in it
	if numPeriods <= 1 {
		return in
	}

	// replace all but the last period with underscore
	return strings.Replace(in, ".", "_", numPeriods-1)
}

func SideEffectName(s SideEffect) string {
	switch s {
	case REQUEST:
		return "request"
	case SUCCESS:
		return "success"
	case FAILURE:
		return "failure"
	case CANCEL:
		return "cancel"
	default:
		return ""
	}
}

func CrudName(crud Crud, repeated bool) string {
	switch crud {
	case CREATE:
		return "create"
	case GET:
		{
			if repeated {
				return "list"
			} else {
				return "get"
			}
		}
	case UPDATE:
		return "update"
	case DELETE:
		return "delete"
	case CUSTOM:
		return "custom"
	default:
		return ""
	}
}

func Tabs(n int) string {
	output := ""
	for i := 0; i < n; i++ {
		output += "  "
	}
	return output
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsFile(s []*gp.FileDescriptorProto, f *gp.FileDescriptorProto) bool {
	for _, a := range s {
		if a.GetName() == f.GetName() {
			return true
		}
	}
	return false
}

func GetFilePath(name string) string {
	// replace proto filetype
	name = strings.Replace(name, ".proto", "_pb", 1)
	// name = name[:strings.LastIndex(name, "/")]
	return name
}

func CreateAggregateByPackage(msgFiles []*gp.FileDescriptorProto, protocTsPath string, statePkg string) []*File {
	var packageNames []string
	var result string
	out := make([]*File, 0)

	for _, file := range msgFiles {
		if !contains(packageNames, file.GetPackage()) && statePkg != file.GetPackage() {
			result = ""
			packageNames = append(packageNames, file.GetPackage())

			for _, f := range msgFiles {
				if f.GetPackage() == file.GetPackage() {
					fp := GetFilePath(f.GetName())
					packageUnderscore := strings.Replace(file.GetPackage(), ".", "/", -1)
					index := strings.LastIndex(fp, "/") + 1
					result += fmt.Sprintf("export * from \"%s%s/%s\";\n", protocTsPath, packageUnderscore, fp[index:])
				}
			}

			out = append(out, &File{
				// replace package name periods with underscores
				Name:    fmt.Sprintf("%s_aggregate.ts", strings.Replace(file.GetPackage(), ".", "_", -1)),
				Content: result,
			})
		}
	}
	return out
}

// find a method descriptor from the annotation string name
func FindMethodDescriptor(serviceFiles []*gp.FileDescriptorProto, fullMethodName string) (*gp.MethodDescriptorProto, error) {
	for _, servFile := range serviceFiles {
		packageName := servFile.GetPackage()
		for _, service := range servFile.GetService() {
			serviceName := service.GetName()
			for _, method := range service.GetMethod() {
				if fmt.Sprintf("%s.%s.%s", packageName, serviceName, method.GetName()) == fullMethodName {
					// make sure it doesn't use client-side streaming (not supported with grpc-web)
					if method.GetClientStreaming() {
						return nil, fmt.Errorf("Client-side streaming not supported. Failed on method: %s", fullMethodName)
					}
					return method, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("Unable to locate method: \"%s\". Missing Method Declaration in Service.", fullMethodName)
}

func FindDescriptor(protos []*gp.FileDescriptorProto, fullMessageName string) (*gp.DescriptorProto, *gp.FileDescriptorProto, string, error) {
	for _, file := range protos {
		packageName := file.GetPackage()
		for _, message := range file.GetMessageType() {
			msgName := fmt.Sprintf(".%s.%s", packageName, message.GetName())
			if msgName == fullMessageName {
				return message, file, "", nil
			}

			// check nested types too
			nested := message.GetNestedType()
			win, desc, depth := checkNestedType(msgName, nested, fullMessageName)
			if win {
				return desc, file, depth, nil
			}
		}
	}
	return nil, nil, "", fmt.Errorf("Unable to locate message: \"%s\". Perhaps the file wasn't listed as a dependency?", fullMessageName)
}

func checkNestedType(prefix string, nested []*gp.DescriptorProto, goal string) (bool, *gp.DescriptorProto, string) {
	for _, n := range nested {
		// full name of the object
		nestedName := fmt.Sprintf("%s.%s", prefix, n.GetName())
		// check for the match
		if nestedName == goal {
			return true, n, prefix
		}
		// if it has nested types, recursively call this function with the updated name
		if len(n.GetNestedType()) != 0 {
			// return checkNestedType(nestedName+"."+n.GetName(), n.GetNestedType(), goal)
			win, desc, depth := checkNestedType(nestedName, n.GetNestedType(), goal)
			if win {
				return true, desc, depth
			}
		}
	}
	// break case
	return false, nil, ""
}

func FindParentMessage(possible_parent *gp.DescriptorProto, possible_child *gp.DescriptorProto) (bool, *gp.DescriptorProto) {
	nested := possible_parent.GetNestedType()

	for _, n := range nested {
		// check for the match
		if n.GetName() == possible_child.GetName() {
			return true, possible_parent
		}

		// if it has nested types, recursively call this function with the updated name
		if len(n.GetNestedType()) != 0 {
			found, message := FindParentMessage(n, possible_child)
			if found {
				return true, message
			}
		}
	}
	// break case
	return false, nil
}
