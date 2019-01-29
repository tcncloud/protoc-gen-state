package redux3

const ToMessageImportTemplate = `{{range $i, $e := .}}
import {{$e.ModuleName}} from '{{$e.FileName}}';{{end}}`

const ToMessageMappingTemplate = `
const messageMap = new Map();
{{range $i, $e := .}}
const {{$e.MapName}} = new Map();{{range $x, $t := $e.TypeLines }}
{{$e.MapName}}.set('{{$t.Name}}', {{$t.FileName}}{{$t.TypeName}});{{end}}
messageMap.set({{$e.FileName}}.{{$e.TypeName}}, {{$e.MapName}});
{{end}}`

const ToMessageTemplate = `

function getNestedMessageConstructor(messageType, fieldName) {
  return messageMap.has(messageType) && messageMap.get(messageType).get(fieldName);
}

export function toMessage(obj: any, messageClass: any) {
  if (!obj) {
    return new messageClass();
  }
  {{if .}}console.groupCollapsed('toMessage');{{end}}
  const message = new messageClass();

  Object.keys(obj).forEach(key => {
    {{if .}}console.groupCollapsed('field:', key);{{end}}
    let ele = obj[key];
    const upperCaseKey = key.charAt(0).toUpperCase() + key.substr(1);
    const setterName = "set" + upperCaseKey;
    const getterName = "get" + upperCaseKey;
    {{if .}}console.log('ele:', ele);
    console.log('typeof ele:', typeof ele);
    console.log('setterName:', setterName);
    console.log('getterName:', getterName);{{end}}

    if (message[setterName]) {
      var nestedMessageContructor = getNestedMessageConstructor(messageClass, key);
      if (nestedMessageContructor) {
        if (key.length > 4 && key.slice(key.length - 4) === 'List' && Array.isArray(ele)) { // check if field is repeated
          {{if .}}console.log('REPEATED field');{{end}}
          ele = ele.map(subEle => toMessage(subEle, nestedMessageContructor));
        } else {
          {{if .}}console.log('regular field');{{end}}
          ele = toMessage(ele, nestedMessageContructor);
        }
      }

      message[setterName](ele);
    } else if (message[getterName] && key.slice(key.length - 3) === 'Map') { // check if field is a map
      {{if .}}console.log('MAP field');{{end}}
      // if the map field is missing, nothing needs to be done.
      if (ele !== undefined && ele !== null) {
        if (Array.isArray(ele)) {
          if (ele.length) {
            var mapObj = message[getterName]();
            var mappedFieldValueConstructor = getNestedMessageConstructor(messageClass, key);
            if (mappedFieldValueConstructor) {
              {{if .}}console.groupCollapsed('keys & values, unserialized');{{end}}
              ele = ele.map(([key, value]) => {
                {{if .}}console.log('key:', key);
                console.log('value:', value);{{end}}
                return [key, mappedFieldValueConstructor(value)];
              });
              {{if .}}console.groupEnd();{{end}}
            }
            {{if .}}console.groupCollapsed('keys & values, serialized');{{end}}
            ele.forEach(([key, value]) => {
              {{if .}}console.log('key:', key);
              console.log('value:', value);{{end}}
              mapObj.set(key, value);
            });
            {{if .}}console.groupEnd();{{end}}
          }
        } else {
          {{if .}}console.groupEnd();
          console.groupEnd();{{end}}
          throw new Error("Protoc-gen-state: Expected field " + key + " to be an array of tuples.");
        }
      }
    } else {
      {{if .}}console.groupEnd();
      console.groupEnd();{{end}}
      throw new Error("No corresponding gRPC setter method for given key: " + key);
    }
    {{if .}}console.groupEnd();{{end}}
  });
  {{if .}}console.groupEnd();{{end}}

  return message;
}
`
