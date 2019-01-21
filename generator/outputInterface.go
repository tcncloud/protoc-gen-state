package generator

import (
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/tcncloud/protoc-gen-state/generator/outputs/redux3"
	"github.com/tcncloud/protoc-gen-state/generator/outputs/redux4"
	"github.com/tcncloud/protoc-gen-state/state"
)

type Outputter interface {
	SetOutputType(state.OutputTypes)
	CreateStateFile([]*gp.FieldDescriptorProto, bool) (*File, error)
	CreateReducerFile([]*gp.FieldDescriptorProto, bool) (*File, error)
	CreateEpicFile([]*gp.FieldDescriptorProto, []*gp.FieldDescriptorProto, []*gp.FileDescriptorProto, int64, int64, string, string, string, int64, int64, bool) (*File, error)
	CreateActionFile([]*gp.FieldDescriptorProto, []*gp.FieldDescriptorProto, []*gp.FileDescriptorProto, bool) (*File, error)
	CreateToMessageFile([]*gp.FileDescriptorProto, []*gp.FileDescriptorProto, string, bool) (*File, error)
	CreateAggregateTypesFile([]*gp.FileDescriptorProto, string) (*File, error)
	CreateAggregateServicesFile([]*gp.FileDescriptorProto, string, string) (*File, error)
}

type StateFileType struct {
	Template string
}
type ReducerFileType struct {
	Template string
}
type ActionFileType struct {
	ImportTemplate string
	GetTemplate    string
	ListTemplate   string
	ResetTemplate  string
	CreateTemplate string
	DeleteTemplate string
	UpdateTemplate string
	CustomTemplate string
}
type EpicFileType struct {
	Template string
}
type ToMessageFileType struct {
	ImportTemplate  string
	MappingTemplate string
	Template        string
}
type AggregatorFileType struct {
	TypeTemplate    string
	ServiceTemplate string
	ExportsTemplate string
}

// The generic outputter will output using all the same methods we have been using, but if an output came along needing to be slightly different (maybe the entity structure needs to change), they can just create a struct that still satisfies the outputter interface
type GenericOutputter struct {
	OutputType     state.OutputTypes
	StateFile      *StateFileType
	ReducerFile    *ReducerFileType
	ActionFile     *ActionFileType
	EpicFile       *EpicFileType
	ToMessageFile  *ToMessageFileType
	AggregatorFile *AggregatorFileType
}

func (this *GenericOutputter) SetOutputType(outputType state.OutputTypes) {
	switch state.OutputTypes_name[int32(outputType)] {
	case "redux4":
		this.StateFile.Template = redux4.StateTemplate
		this.ReducerFile.Template = redux4.ReducerTemplate
		this.ActionFile.ImportTemplate = redux4.ActionImportTemplate
		this.ActionFile.GetTemplate = redux4.ActionGetTemplate
		this.ActionFile.ListTemplate = redux4.ActionListTemplate
		this.ActionFile.ResetTemplate = redux4.ActionResetTemplate
		this.ActionFile.CreateTemplate = redux4.ActionCreateTemplate
		this.ActionFile.DeleteTemplate = redux4.ActionDeleteTemplate
		this.ActionFile.UpdateTemplate = redux4.ActionUpdateTemplate
		this.ActionFile.CustomTemplate = redux4.ActionCustomTemplate
		this.EpicFile.Template = redux4.EpicTemplate
		this.ToMessageFile.ImportTemplate = redux4.ToMessageImportTemplate
		this.ToMessageFile.MappingTemplate = redux4.ToMessageMappingTemplate
		this.ToMessageFile.Template = redux4.ToMessageTemplate
		this.AggregatorFile.TypeTemplate = redux4.AggregatorTypeTemplate
		this.AggregatorFile.ServiceTemplate = redux4.AggregatorServiceTemplate
		this.AggregatorFile.ExportsTemplate = redux4.AggregatorExportsTemplate
	default: // defaults to redux3
		this.StateFile.Template = redux3.StateTemplate
		this.ReducerFile.Template = redux3.ReducerTemplate
		this.ActionFile.ImportTemplate = redux3.ActionImportTemplate
		this.ActionFile.GetTemplate = redux3.ActionGetTemplate
		this.ActionFile.ListTemplate = redux3.ActionListTemplate
		this.ActionFile.ResetTemplate = redux3.ActionResetTemplate
		this.ActionFile.CreateTemplate = redux3.ActionCreateTemplate
		this.ActionFile.DeleteTemplate = redux3.ActionDeleteTemplate
		this.ActionFile.UpdateTemplate = redux3.ActionUpdateTemplate
		this.ActionFile.CustomTemplate = redux3.ActionCustomTemplate
		this.EpicFile.Template = redux3.EpicTemplate
		this.ToMessageFile.ImportTemplate = redux3.ToMessageImportTemplate
		this.ToMessageFile.MappingTemplate = redux3.ToMessageMappingTemplate
		this.ToMessageFile.Template = redux3.ToMessageTemplate
		this.AggregatorFile.TypeTemplate = redux3.AggregatorTypeTemplate
		this.AggregatorFile.ServiceTemplate = redux3.AggregatorServiceTemplate
		this.AggregatorFile.ExportsTemplate = redux3.AggregatorExportsTemplate
	}
}

func MakeGenericOutputter(outputType state.OutputTypes) *GenericOutputter {
	outputter := &GenericOutputter{
		OutputType:     outputType,
		StateFile:      &StateFileType{},
		ReducerFile:    &ReducerFileType{},
		ActionFile:     &ActionFileType{},
		EpicFile:       &EpicFileType{},
		ToMessageFile:  &ToMessageFileType{},
		AggregatorFile: &AggregatorFileType{},
	}
	outputter.SetOutputType(outputType)
	return outputter
}

// An outputter might need to be customized more than the generic use case.
type MobxOutputter struct{}

func (this *MobxOutputter) SetOutputType(outputType state.OutputTypes) {
	// do something unique
}
func (this *MobxOutputter) CreateStateFile(stateFields []*gp.FieldDescriptorProto, debug bool) (*File, error) {
	// do something unique
	return &File{
		Name:    "state_pb.ts",
		Content: "Not implemented",
	}, nil
}
func (this *MobxOutputter) CreateActionFile(stateFields []*gp.FieldDescriptorProto, customFields []*gp.FieldDescriptorProto, serviceFiles []*gp.FileDescriptorProto, debug bool) (*File, error) {
	return &File{
		Name:    "actions_pb.ts",
		Content: "Not implemented",
	}, nil
}
func (this *MobxOutputter) CreateReducerFile(stateFields []*gp.FieldDescriptorProto, debug bool) (*File, error) {
	return &File{
		Name:    "reducer_pb.ts",
		Content: "Not implemented",
	}, nil
}
func (this *MobxOutputter) CreateEpicFile(stateFields []*gp.FieldDescriptorProto, customFields []*gp.FieldDescriptorProto, serviceFiles []*gp.FileDescriptorProto, defaultTimeout int64, defaultRetries int64, authTokenLocation string, hostnameLocation string, hostname string, portin int64, debounce int64, debug bool) (*File, error) {
	return &File{
		Name:    "epics_pb.ts",
		Content: "Not implemented",
	}, nil
}
func (this *MobxOutputter) CreateToMessageFile(servFiles []*gp.FileDescriptorProto, protos []*gp.FileDescriptorProto, protocTsPath string, debug bool) (*File, error) {
	return &File{
		Name:    "to_message_pb.ts",
		Content: "Not implemented",
	}, nil
}
func (this *MobxOutputter) CreateAggregateServicesFile(serviceFiles []*gp.FileDescriptorProto, protocTsPath string, statePkg string) (*File, error) {
	return &File{
		Name:    "protoc_services_pb.ts",
		Content: "Not implemented",
	}, nil
}
func (this *MobxOutputter) CreateAggregateTypesFile(msgFiles []*gp.FileDescriptorProto, statePkg string) (*File, error) {
	return &File{
		Name:    "protoc_types_pb.ts",
		Content: "Not implemented",
	}, nil
}
