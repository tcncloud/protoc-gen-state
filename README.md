# protoc-gen-state - WORK IN PROGRESS
Seriously, give it a minute

## internal: test usage and generation
| command | result |
| ------- | ------ |
| `make` | <ul><li>builds the plugin</li><li>runs it on the test proto</li><li>creates an output folder called generated</li><li>runs tests based on those generated files</li></ul> |
| `make clean` | cleans it up (removes binary and generated folder) |


---


### NOTES
##### generate options proto for go
`protoc --go_out=. github.com/tcncloud/protoc-gen-state/state/*.proto` from __$(GOPATH)/src__
