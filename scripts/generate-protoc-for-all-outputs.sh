 #!/bin/bash

#Change subshell working directory to root directory of this repo
cd "$(cd "$(dirname "${BASH_SOURCE[0]}")/../" > /dev/null && pwd)" || return

GENERATED='./generated'

generate_ts() { # first parameter should be the name of the output directory
  protoc --plugin="protoc-gen-ts=node_modules/.bin/protoc-gen-ts" \
    --js_out="import_style=commonjs,binary:." \
    --ts_out="service=true:." \
    e2e/$1/protos/readinglist/readinglist.proto
}

generate_state() { # first parameter should be the name of the output directory
  # output into the generated dir for our go tests
  mkdir -p "$GENERATED/$1"
  protoc -I. -I./e2e/$1/protos -I./state/options.proto \
    --plugin=./protoc-gen-state \
    --state_out=$GENERATED/$1 ./e2e/$1/protos/basic.proto

  # output into the e2e dir for our js tests
  mkdir -p "e2e/$1/protos/BasicState"
  protoc -I. -I./e2e/$1/protos -I./state/options.proto \
    --plugin=./protoc-gen-state \
    --state_out=./e2e/$1/protos/BasicState ./e2e/$1/protos/basic.proto
}


# This command finds the code block of "enum OutputTypes" in the state/options.proto file e.g.:
  # enum OutputTypes {
  #   redux3 = 0;
  #   redux4 = 1;
  # }
# then it gets the first word of each field
for line in $(sed -n -e '/enum OutputTypes {/,/}/ p' state/options.proto | sed '1d;$d' | awk ' { print $1 } ')
do
  echo "generating: $line"
  generate_ts "$line"
  generate_state "$line"
done
