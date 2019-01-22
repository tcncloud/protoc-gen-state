package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// . "github.com/tcncloud/protoc-gen-state"
	"github.com/tcncloud/protoc-gen-state/state"
)

var _ = Describe("Main", func() {
	for outputType := range state.OutputTypes_value {
		path := "generated/" + outputType + "/"

		// these files are generated every the plugin is successfully run
		Describe("Static Generated Files", func() {
			files := []string{
				"actions_pb.ts",
				"epics_pb.ts",
				"protoc_services.ts",
				"protoc_types.ts",
				"reducer_pb.ts",
				"state_pb.ts",
				"to_message_pb.ts",
			}

			for _, f := range files {
				It("Should generate the file: "+f, func() {
					Expect(path + f).To(BeAnExistingFile())
				})
			}
		})

		// dynamically creates aggregate files by package name
		Describe("Dynamic Generated Files", func() {
			fileName := "readinglist_aggregate.ts"

			It("Should generate an aggregate file for the readinglist package", func() {
				Expect(path + fileName).To(BeAnExistingFile())
			})
		})
	}
})
