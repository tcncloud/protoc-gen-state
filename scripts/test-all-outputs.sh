 #!/usr/bin/env bash

#Change subshell working directory to root directory of this repo
cd "$(cd "$(dirname "${BASH_SOURCE[0]}")/../" > /dev/null && pwd)" || return

# run go tests:
ginkgo .

# now run js tests
yarn
npx tsc -p "./tsconfig.json"
  # This command finds the code block of "enum OutputTypes" in the state/options.proto file e.g.:
    # enum OutputTypes {
    #   redux3 = 0;
    #   redux4 = 1;
    #   mobx = 2;
    # }
  # then it gets the first word of each field
for line in $(sed -n -e '/enum OutputTypes {/,/}/ p' state/options.proto | sed '1d;$d' | awk ' { print $1 } ')
do
  JASMINE_CONFIG_PATH=e2e/$line/spec/support/jasmine.json node_modules/.bin/ts-node --node_options=--preserve-symlinks --node_options=--max-old-space-size=8192 node_modules/jasmine/bin/jasmine
done
