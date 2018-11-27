# WORK IN PROGRESS 

## protoc-gen-state
Generates redux state (actions, epics, reducer) from a protobuf file. 

Readme is under construction, refer to the example in `e2e/src` for usage.

### internal: test usage and generation
| command | result |
| ------- | ------ |
| `make` | <ul><li>builds the plugin</li><li>runs the go tests and cleans up</li><li>runs the javascript tests and cleans up</li></ul>

---

[NOTES]:   <> (### NOTES)
[comment]: <> (`protoc --go_out=. github.com/tcncloud/protoc-gen-state/state/*.proto` )
