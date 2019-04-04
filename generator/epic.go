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
	"fmt"
	gp "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"strconv"
	"strings"
	"text/template"
)

type EpicEntity struct {
	Name             string
	ProtoInputType   string
	ProtoOutputType  string
	FullMethodName   string
	JsonName         string
	Debounce         int64
	Timeout          int64
	Retries          int64
	Repeat           bool
	Auth             string
	Hostname         string
	HostnameLocation string
	Port             string
	Updater          bool
	Debug            bool
}

func (this *GenericOutputter) CreateEpicFile(stateFields []*gp.FieldDescriptorProto, customFields []*gp.FieldDescriptorProto, serviceFiles []*gp.FileDescriptorProto, defaultTimeout int64, defaultRetries int64, authTokenLocation string, hostnameLocation string, hostname string, portin int64, debounce int64, debug bool) (*File, error) {
	epicEntities := []*EpicEntity{}

	port := ":" + strconv.FormatInt(portin, 10)

	// transform stateFields into our EpicEntity implementation so template can read values
	for _, field := range stateFields {
		repeated := field.GetLabel() == 3

		// verify the method annotations
		fieldAnnotations, err := GetFieldOptions(field)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level annotations: %v", err)
		}

		// field level overrides for timeout/retry
		timeout := fieldAnnotations.GetTimeout()
		if timeout == 0 { // if it wasn't overriden
			timeout = defaultTimeout
		}
		retries := fieldAnnotations.GetRetries()
		if retries == 0 { // if it wasn't overriden
			retries = defaultRetries
		}

		var meth *gp.MethodDescriptorProto
		// get the method for each crud
		for c := CREATE; c < CRUD_MAX; c++ {
			// clear for the loop
			meth = nil

			crudAnnotation := GetAnnotation(*fieldAnnotations.GetMethod(), c, repeated)
			if crudAnnotation != "" {
				meth, err = FindMethodDescriptor(serviceFiles, crudAnnotation)
				if err != nil {
					return nil, err
				}
			}

			if meth != nil {
				// only returns arrays on these
				var repeatEntity bool
				if CrudName(c, repeated) == "list" {
					repeatEntity = true
				} else {
					repeatEntity = false
				}

				var updater bool
				if CrudName(c, repeated) == "update" && repeated {
					updater = true
				} else {
					updater = false
				}

				epicEntities = append(epicEntities, &EpicEntity{
					Name:             CrudName(c, repeated) + strings.Title(*field.JsonName),
					ProtoInputType:   fmt.Sprintf("ProtocTypes.%s", CreatePackageAndTypeString(meth.GetInputType())),
					ProtoOutputType:  fmt.Sprintf("ProtocTypes.%s", CreatePackageAndTypeString(meth.GetOutputType())),
					FullMethodName:   fmt.Sprintf("ProtocServices.%s", FullMethodNameFormat(crudAnnotation)),
					JsonName:         *field.JsonName,
					Debounce:         debounce,
					Timeout:          timeout,
					Retries:          retries,
					Repeat:           repeatEntity,
					Auth:             authTokenLocation,
					Hostname:         hostname,
					HostnameLocation: hostnameLocation,
					Port:             port,
					Updater:          updater,
					Debug:            debug,
				})
			}
		}
	}

	// do the same for customActions
	// TODO combine the logic
	for _, field := range customFields {
		repeated := field.GetLabel() == 3

		// verify the method annotations
		fieldAnnotations, err := GetFieldOptions(field)
		if err != nil {
			return nil, fmt.Errorf("Error getting field level annotations: %v", err)
		}

		timeout := fieldAnnotations.GetTimeout()
		if timeout == 0 { // if it wasn't overriden
			timeout = defaultTimeout
		}
		retries := fieldAnnotations.GetRetries()
		if retries == 0 { // if it wasn't overriden
			retries = defaultRetries
		}

		var meth *gp.MethodDescriptorProto

		crudAnnotation := fieldAnnotations.GetMethod().GetCustom()
		if crudAnnotation != "" {
			meth, err = FindMethodDescriptor(serviceFiles, crudAnnotation)
			if err != nil {
				return nil, err
			}
		}

		if meth != nil {

			// TODO uses repeated from the field name, should use the output type
			epicEntities = append(epicEntities, &EpicEntity{
				Name:             "custom" + strings.Title(*field.JsonName),
				ProtoInputType:   fmt.Sprintf("ProtocTypes.%s", CreatePackageAndTypeString(meth.GetInputType())),
				ProtoOutputType:  fmt.Sprintf("ProtocTypes.%s", CreatePackageAndTypeString(meth.GetOutputType())),
				FullMethodName:   fmt.Sprintf("ProtocServices.%s", FullMethodNameFormat(crudAnnotation)),
				JsonName:         *field.JsonName,
				Debounce:         debounce,
				Timeout:          timeout,
				Retries:          retries,
				Repeat:           repeated,
				Auth:             authTokenLocation,
				Hostname:         hostname,
				HostnameLocation: hostnameLocation,
				Port:             port,
				Debug:            debug,
			})
		}
	}

	tmpl := template.Must(template.New("epic").Parse(this.EpicFile.Template))

	var output bytes.Buffer
	tmpl.Execute(&output, epicEntities)

	return &File{
		Name:    "epics_pb.ts",
		Content: output.String(),
	}, nil
}

func FullMethodNameFormat(name string) string {
	index := strings.LastIndex(name, ".")        // first
	index = strings.LastIndex(name[:index], ".") // second
	return strings.Replace(name, name[:index], strings.Replace(name[:index], ".", "_", -1), 1)
}
