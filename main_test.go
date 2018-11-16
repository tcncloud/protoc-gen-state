package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// . "github.com/tcncloud/protoc-gen-state"
)

var _ = Describe("Main", func() {
	Describe("Generates Files", func() {
		path := "generated/"
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

})
