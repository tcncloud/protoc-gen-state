 #!/usr/bin/env bash

#Change subshell working directory to root directory of this repo
cd "$(cd "$(dirname "${BASH_SOURCE[0]}")/../" > /dev/null && pwd)" || return

GENERATED='./generated'


rm -rf node_modules/
rm -rf $GENERATED


# This command finds the code block of "enum OutputTypes" in the state/options.proto file e.g.:
# enum OutputTypes {
#   redux3 = 0;
#   redux4 = 1;
# }
# then it gets the first word of each field
for line in "$(sed -n -e '/enum OutputTypes {/,/}/ p' state/options.proto | sed '1d;$d' | awk ' { print $1 } ')"
do
  rm -rf e2e/$line/protos/BasicState
  rm -rf e2e/$line/protos/readinglist/{*.ts,*.js}
done
