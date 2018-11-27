# Copyright 2017, TCN Inc.
# All rights reserved.

# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions are
# met:

#     * Redistributions of source code must retain the above copyright
# notice, this list of conditions and the following disclaimer.
#     * Redistributions in binary form must reproduce the above
# copyright notice, this list of conditions and the following disclaimer
# in the documentation and/or other materials provided with the
# distribution.
#     * Neither the name of TCN Inc. nor the names of its
# contributors may be used to endorse or promote products derived from
# this software without specific prior written permission.

# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
# "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
# LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
# A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
# OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
# SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
# LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
# DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
# THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
# (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

GENERATED=generated/

all: deps build test

build: gen protos

test: test-go clean-go test-js clean-js

gen:
	go build .

protos:
	mkdir -p $(GENERATED)
	protoc -I. -I./e2e/protos -I./state/options.proto \
		--plugin=./protoc-gen-state \
		--state_out=$(GENERATED) ./e2e/protos/basic.proto

deps:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/iancoleman/strcase

clean-go:
	# rm -f ./protoc-gen-state
	rm -rf $(GENERATED)

clean-js:
	rm -rf node_modules/
	rm -rf e2e/generated/

test-go:
	ginkgo .

test-js:
	yarn
	yarn run build
	npx tsc -p "./tsconfig.json"
	yarn run test

# test - generate multiple proto files, panic
# test - <1 or >3 messages in state proto, panic
# test - messages besides ReduxState, CustomActions, ExternalLink in state proto, panic
