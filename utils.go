package main

import (
	"fmt"
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"strings"
)

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
	// grab filename without path
	// replace proto filetype
	name = strings.Replace(name, ".proto", "_pb", 1)
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
					result += fmt.Sprintf("export * from \"%s%s\";\n", protocTsPath, GetFilePath(f.GetName()))
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
