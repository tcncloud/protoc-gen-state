package main

import (
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func CreateStateFile(stateFields []*gp.FieldDescriptorProto) (*File, error) {
	// TODO
	return &File{
		Name:    "state_pb.ts",
		Content: "placeholder",
	}, nil
}
