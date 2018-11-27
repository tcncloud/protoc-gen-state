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
	"bytes"
	"strings"
	"text/template"

	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

const typeAggregate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

{{range $i, $e := .}}
import * as {{$e.Package}} from "./{{$e.Package}}_aggregate";{{end}}
{{range $i, $e := .}}
export { {{$e.Package}} };{{end}}`

type TypeEntity struct {
	Package string
}

func CreateAggregateTypesFile(msgFiles []*gp.FileDescriptorProto, statePkg string) (*File, error) {
	typeEntities := []*TypeEntity{}
	packageNames := []string{statePkg}

	for _, file := range msgFiles {
		if !contains(packageNames, file.GetPackage()) {
			packageNames = append(packageNames, file.GetPackage())
			typeEntities = append(typeEntities, &TypeEntity{Package: file.GetPackage()})
		}
	}

	tmpl := template.Must(template.New("types").Parse(typeAggregate))
	var output bytes.Buffer
	tmpl.Execute(&output, typeEntities)

	return &File{
		Name:    "protoc_types_pb.ts",
		Content: output.String(),
	}, nil
}

const serviceAggregate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

{{range $i, $e := .}}
import * as {{$e.Name}}_service_in from "{{$e.Location}}_service";{{end}}`

const serviceExports = `
{{range $i, $e := .}}
export var {{$e.Package}} = { {{range $j, $x := $e.Exports}}
	...{{$x}}_service_in,{{end}}
}{{end}}`

type ServiceEntity struct {
	Location string
	Name     string
	Package  string
}
type ServiceExport struct {
	Package string
	Exports []string
}

func CreateAggregateServicesFile(serviceFiles []*gp.FileDescriptorProto, protocTsPath string, statePkg string) (*File, error) {
	serviceEntities := []*ServiceEntity{}
	exportEntities := []*ServiceExport{}
	packageNames := []string{statePkg}

	for f, file := range serviceFiles {
		if !contains(packageNames, file.GetPackage()) {
			packageNames = append(packageNames, file.GetPackage())

			exports := []string{}
			for i := f; i < len(serviceFiles); i++ {
				if serviceFiles[i].GetPackage() == file.GetPackage() {
					filePathOriginal := GetFilePath(serviceFiles[i].GetName())
					index := strings.LastIndex(filePathOriginal, "/") + 1
					filePath := filePathOriginal[index:]
					name := strings.Replace(strings.ToLower(filePathOriginal), "/", "_", -1)
					exports = append(exports, name)
					serviceEntities = append(serviceEntities, &ServiceEntity{
						Location: protocTsPath + filePath,
						Name:     name,
						Package:  file.GetPackage(),
					})
				}
			}

			exportEntities = append(exportEntities, &ServiceExport{
				Package: file.GetPackage(),
				Exports: exports,
			})
		}
	}

	tmpl := template.Must(template.New("services").Parse(serviceAggregate))
	exp := template.Must(template.New("exports").Parse(serviceExports))
	var output bytes.Buffer
	tmpl.Execute(&output, serviceEntities)
	exp.Execute(&output, exportEntities)

	return &File{
		Name:    "protoc_services_pb.ts",
		Content: output.String(),
	}, nil
}
