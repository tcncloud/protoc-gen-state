 #!/usr/bin/env bash

#Change subshell working directory to root directory of this repo
cd "$(cd "$(dirname "${BASH_SOURCE[0]}")/../" > /dev/null && pwd)" || return

GENERATED='./generated'

generate_protos() {
  mkdir -p "$GENERATED/$1"

  protoc -I. -I./e2e/$1/protos -I./state/options.proto \
    --plugin=./protoc-gen-state \
    --state_out=$GENERATED ./e2e/$1/protos/basic.proto

  cp ./protoc-gen-state ./e2e/$1
}

while read line; do
  generate_protos "$line"
done <./outputs.txt
  
