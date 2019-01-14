package generator

import (
  "github.com/tcncloud/protoc-gen-state/generator/redux3"
  "github.com/tcncloud/protoc-gen-state/generator/redux4"
  "github.com/tcncloud/protoc-gen-state/state"
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Outputter interface {
  CreateStateFile([]*gp.FieldDescriptorProto, state.OutputTypes, bool) (*File, error)
  // CreateEpicFile([]*gp.FieldDescriptorProto, state.OutputTypes, []*gp.FieldDescriptorProto, []*gp.FileDescriptorProto, int64, int64, string, string, string, int64, int64, bool) (*File, error)
}


type StateFileType struct {
  // Entities []StateEntity
  Template string
}

// The generic outputter will output using all the same methods we have been using, but if an output came along needing to be slightly different (maybe the entity structure needs to change), they can just create a struct that still satisfies the outputter interface
type GenericOutputter struct {
  OutputType state.OutputTypes
  StateFile *StateFileType
}


func (this *GenericOutputter) SetOutputType(outputType state.OutputTypes) {
  switch state.OutputTypes_name[int32(outputType)] {
  case "redux3":
    this.StateFile.Template = redux3.StateTemplate
  case "redux4":
    this.StateFile.Template = redux4.StateTemplate
  default:
    this.StateFile.Template = redux3.StateTemplate
  }
}

