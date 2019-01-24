 #!/usr/bin/env bash

#Change subshell working directory to root directory of this repo
cd "$(cd "$(dirname "$0")/../" > /dev/null && pwd)" || return

# run go tests:
ginkgo .

# now run js tests
  # This command finds the code block of "enum OutputTypes" in the state/options.proto file e.g.:
    # enum OutputTypes {
    #   redux3 = 0;
    #   redux4 = 1;
    #   mobx = 2;
    # }
  # then it gets the first word of each field
for line in $(sed -n -e '/enum OutputTypes {/,/}/ p' state/options.proto | sed '1d;$d' | awk ' { print $1 } ')
do
  echo "TESTING $line"
  cd e2e/$line
  yarn

  yarn run test
  cd ../../
done
