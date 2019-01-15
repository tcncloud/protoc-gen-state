package redux3
       
const AggregatorTypeTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */
/* typeAggregate */

{{range $i, $e := .}}
import * as {{$e.Package}} from "./{{$e.Package}}_aggregate";{{end}}
{{range $i, $e := .}}
export { {{$e.Package}} };{{end}}`

const AggregatorServiceTemplate = `/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */
/* serviceAggregate */

{{range $i, $e := .}}
import * as {{$e.Name}}_service_in from "{{$e.Location}}_service";{{end}}`

const AggregatorExportsTemplate = `
{{range $i, $e := .}}
export var {{$e.Package}} = {
{{range $j, $x := $e.Exports}}  ...{{$x}}_service_in,{{end}}
}{{end}}`

