# protoc-gen-state - WORK IN PROGRESS
Seriously, give it a minute

## internal: test usage and generation
`make`       - should create a folder called `generated` with the output 

`make clean` - cleans it up (removes binary and generated folder)


---


### NOTES
##### generate options proto for go
`protoc --go_out=. github.com/tcncloud/protoc-gen-state/state/*.proto` from __$(GOPATH)/src__
