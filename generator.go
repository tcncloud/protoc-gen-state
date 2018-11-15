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
	// "fmt"
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"strings"
)

func generate(filepaths []string, protos []*gp.FileDescriptorProto) ([]*File, error, error) {
	// file descriptor proto slice
	oracle := Oracle{protos}
	// list of output files
	out := make([]*File, 0)
	// get struct of { package, go_package }
	pkgs := oracle.Packages()

	for _, pkg := range pkgs {
		// get all files that match the package name
		files := oracle.GenerationFilesIn(&pkg)
		if len(files) == 0 {
			continue
		}
		// append to the output but change the name of the file
		out = append(out, &File{
			Name:    strings.Replace(pkg.Name, ".", "/", -1) + "/" + pkg.Name + ".generated.proto",
			Content: "test",
		})
	}
	return out, nil, nil
}

type Oracle struct {
	protos []*gp.FileDescriptorProto
}

// returns true if the package name is a dependency for any of the files in oracle
func (o Oracle) IsDependency(name string) bool {
	for _, f := range o.protos {
		for _, d := range f.Dependency {
			if name == d {
				return true
			}
		}
	}
	return false
}

// for every file in the oracle slice,
// make a struct of { package, go_package }
func (o Oracle) Packages() []Package {
	pkgs := make(map[Package]struct{})

	for _, f := range o.protos {
		pkgs[Package{
			Name: f.GetPackage(),
			GoPkg: func() string {
				if opts := f.GetOptions(); opts != nil {
					return opts.GetGoPackage()
				}
				return ""
			}(),
		}] = struct{}{}
	}
	out := make([]Package, 0)

	for p, _ := range pkgs {
		out = append(out, p)
	}
	return out
}

// for now just gathers all files in that package name
func (o Oracle) GenerationFilesIn(pkg *Package) []*gp.FileDescriptorProto {
	// gather the results
	out := make([]*gp.FileDescriptorProto, 0)
	// return all files that match this package name
	files := o.FilesIn(pkg)
	// add all the files to the output
	for _, f := range files {
		// exclude deps like google/protobuf
		if o.IsDependency(f.GetName()) {
			continue
		}
		out = append(out, f)
	}
	return out
}

// subset of oracle that matches the package name (not go_package)
func (o Oracle) FilesIn(p *Package) []*gp.FileDescriptorProto {
	var out []*gp.FileDescriptorProto
	for _, f := range o.protos {
		if f.GetPackage() == p.Name {
			out = append(out, f)
		}
	}
	return out
}

type File struct {
	Name    string
	Content string
}

type Package struct {
	Name  string // name of the protobuf package given to the descriptor
	GoPkg string // the go_package option if there is one
}
