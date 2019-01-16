# VERY ALPHA CODE !!!!
## protoc-gen-state

[![Build Status](https://www.travis-ci.org/tcncloud/protoc-gen-state.svg?branch=master)](https://www.travis-ci.org/tcncloud/protoc-gen-state)

### General Overview

This is a protoc plugin that generates redux state (actions, reducer, epics, state) to match a State protobuf message. If your app only has CRUD operations for data types, it's as simple as dispatching an action and retrieving the value from redux. Can also handle non-CRUD actions and epics for more generic api calls.

Run the tests with `yarn test`

### Broad Strokes

Define a [.proto](#example) file with three messages.
 * **ReduxState** - The only required message. Will generate redux state based on the fields.
   - [Actions](#actions_pb): Will generate actions for each field with the [CLUDG](https://cloud.google.com/apis/design/standard_methods) method prepended and Request, Success, Failure, or Cancel appended. Will also generate an action to reset the state that has no payload. A resolve and reject callback can optionally be passed to Request actions and resolved as a promise. (The epics will call the callback before dispatching success/failure actions)
     * [create|update|delete|get|list]**Name**[Request|RequestPromise|Success|Failure|Cancel]
     * reset**Name**
   - [Reducer](#reducer_pb): Will generate reducer blocks to handle the side-effects of each action accordingly. CustomActions will not generate reducer blocks. 
   - [Epics](#epics_pb): Will generate an epic that listens for a Request action, sends an api call, and dispatches a Success or Failure action depending on the result. If a resolve and reject callback was provided, it will call those before dispatching.

 * **CustomActions** - Optional. Will define actions and epics for non-CLUDG generic methods. Should be handled with a custom reducer or simply by calling the RequestPromise action.
   - [Actions](#actions_pb): Only the method annotion `custom` is allowed so it will prepend actions with Custom and append them with Request, Success, Failure, and Cancel the same as the actions above.
     * custom**Name**[Request|RequestPromise|Success|Failure|Cancel]

### Annotations
Configuration for the plugin is provided by annotations inside the proto file.

| required | level | name | description | default |
| -------- | ----- | ---- | ----------- | ------- |
|x| message | state_options | defines a message as REDUX_STATE (generates reducer/epic/action) or CUSTOM_ACTION (generates only epic/actions) | |
|x| file | hostname | set a static hostname string | |
|x| file | hostname_location | defines a location in redux to grab a dynamic hostname | |
|x| file | protoc_ts_path | defines the path to generated typescript files | |
|x| file | root_path | defines the path to the rootAction file inside your app | |
| | file | auth_token_location | defines the location in redux of an auth token to pass with the Authorization Bearer header in api calls, otherwise no header | |
| | file | debounce | sets the debounce for all api calls | 400 |
| | file | debug | adds helpful console logs generated files | false |
| | file | default_retries | sets the file level default number of retries | 0 |
| | file | default_timeout | sets the file level default timeout in ms | 15000 |
| | file | port | port to append after the hostname for api calls | |
| | field | retries | overrides the default retries for a specific field | |
| | field | timeout | overrides the default timeout for a specific [CLUDG](https://cloud.google.com/apis/design/standard_methods) method on a specific field | |
|x| field.method | method.\<cludg\> | defines the full RPC method name (package.service.method) to be used for a specific [CLUDG](https://cloud.google.com/apis/design/standard_methods) method for a specific field | |
|x| field.method | method.custom | defines the full RPC method name (package.service.method) for a custom action. Required for each CustomAction. Throws an error if used elsewhere. | |
| | field.method | method_timeout.\<cludg\> | overrides all timeouts for this specific [CLUDG](https://cloud.google.com/apis/design/standard_methods) method | |
| | field.method | method_retries.\<cludg\> | overrides all retries for this specific [CLUDG](https://cloud.google.com/apis/design/standard_methods) method | |

Notes:
 * one hostname annotation is required, if neither or both are provided an error will be thrown.
 * while the annotation `(method_timeout).custom` is allowed, it is better to simply write `(timeout)` since there is only one method allowed on custom actions.

### Example
Here is an example of a state proto file with all the bells and whistles. The configuration options are defined by the file level annotations at the top.

```
option (state_file_options) = {
  auth_token_location = "auth.idToken";            // redux location of token
  debounce = 500;                                  // debounce timer set to .5 s
  debug = true;                                    // debug console logs in generated code
  default_retries = 0;                             // overrides plugin default
  default_timeout = 7000;                          // overrides plugin default
  hostname_location = "user.selectedRegion.value"; // redux location of hostname
  port = 9090;                                     // port to append to hostname
  protoc_ts_path = "../../";                       // path to generated typescript files
};

message ReduxState {   // generates the redux state
  option (state_options).type = REDUX_STATE; // mark as redux state message

  api.v0alpha.Organization myOrg = 1 [
    (state_field_options).retries = 3,
    (state_field_options).method.create = "api.v0alpha.Org.CreateOrganization",
    (state_field_options).method.delete = "api.v0alpha.Org.DeleteOrganization",
    (state_field_options).method_timeout.create = 3000,
    (state_field_options).method_retries.delete = 4
  ];
}

message CustomAction {  // generates non-CLUDG actions/epics without a reducer
  option (state_options).type = CUSTOM_ACTION; // mark as custom message, so no reducer generated

  api.v0alpha.Country countryNameChange = 1 [
    (state_field_options).timeout = 4,
    (state_field_options).method.custom = "api.v0alpha.Org.ChangeNameOfCountry"
  ];
}
```

Which will generate the following actions:
* createMyOrgRequest
* createMyOrgRequestPromise // deprecated, use Request with optional resolve/reject callbacks
* createMyOrgSuccess
* createMyOrgFailure
* createMyOrgCancel
* deleteMyOrgRequest
* deleteMyOrgRequestPromise // deprecated, use Request with optional resolve/reject callbacks
* deleteMyOrgSuccess
* deleteMyOrgFailure
* deleteMyOrgCancel
* resetMyOrg
* customCountryNameChangeRequest
* customCountryNameChangeRequestPromise // deprecated, use Request with optional resolve/reject callbacks
* customCountryNameChangeSuccess
* customCountryNameChangeFailure
* customCountryNameChangeCancel

and the following epics:
* createMyOrgEpic
* deleteMyOrgEpic
* customCountryNameChangeEpic

with reducer blocks that handle all of the actions accordingly.

The plugin will match the method annotation string with the corresponding rpc method from the service files. The signature of the action creators will match the rpc method signature. For example, if an rpc returns a stream of objects the Success action will pass an array of objects. A more in-depth description of action files generated can be found: [here](#actions_pb)

--- 

## Granular Overview

### protoc_types_pb
exports all types output from ts-protoc-gen nested by package name (snake case)

### protoc_services_pb 
exports all services and methods output from ts-protoc-gen nested by package name (snake case) followed by service name

### reducer_pb
Imports the types output by ts-protoc-gen, the actions (actions_pb), and the state (state_pb). Creates a reducer that has a switch case for each of the generated actions and mutates the state accordingly.
- {CLUDG}{Element}Request actions set isLoading to true
- {CLUDG}{Element}Cancel actions set isLoading to false
- {CLUDG}{Element}Success actions update the Element, set isLoading to false, and reset the error key
- {CLUDG}{Element}Failure actions set the error field and set isLoading to false

CLUDG actions for repeated elements only modify one element in the array, with the exception of List, which will replace the whole array. Users can still add an action to say, delete the entire array, but this will have to be in its own action and reducer files and will not be inside the generated/state/reducer_pb

Custom actions will not generate reducer cases. Likewise, they will not generate `reset` actions, because those are handled entirely by the reducer.

### state_pb
Imports the types output by ts-protoc-gen. Creates an interface (named ProtocState) defining the State message in TypeScript and an initial state object (named initialState) that matches the type.

```
export interface ProtocState { }
export const initialState : ProtocState { }
```

Each entry from the state proto becomes an object with an isLoading boolean, errorMsg string, and the value itself.

```
myElement: {
  isLoading : boolean;
  errorMsg : string;
  value : types.myElement | null; // types.myElement[] if repeated
}
```

Message types are OR'd with null and are defaulted to null.

Repeated types are not nullable and are defaulted to empty arrays. ([] as Type[])

### epics_pb
Imports the the types (protoc_types_pb), services (protoc_services_pb), and actions (actions_pb). Creates one epic per RPC defined in the protobuf service with the name matching the state message field. They all follow the same general structure:

- Match on {RPC}Request actions
- Convert action payload to {RPC} input (plain object to protobuf message)
- Send {RPC}
- If successful, dispatch {RPC}Success
- Otherwise, dispatch {RPC}Failure
- {RPC}Cancel actions will cancel the api call

note: The epics will call the resolve/reject callbacks before dispatching Success or Failure actions (if provided, they are optional) so that promises can be resolved based on them. (Redux actions don't natively "return" anything)

For example, given this proto:

```
package my.example

service ExampleService {
  rpc CreateBook(Book) returns (Book)
  rpc CreateBorkFromBook(Book) returns (Bork)
}

message Book { ... }
message Bork { ... }
```
```
// generational proto
message ReduxState {
  option (state_options).type = REDUX_STATE;
  my.example.Book myBook = 1 [
    (method).create = "my.example.ExampleService.CreateBook"
  ];
}
message CustomActions {
  option (state_options).type = CUSTOM_ACTION;
  my.example.Bork borkFromBook = 1 [
    (state_field_options).method.custom = "my.example.ExampleService.CreateBorkFromBook"
  ];
}
```

It would generate two epics:
```
createMyBookEpic
customBorkFromBookEpic
```
That listen on Request actions and dispatch Success/Failure actions depending on the success of the grpc call.

### actions_pb

Imports the types output by ts-protoc-gen. Creates three action creators (Request, Success, Failure, Cancel) for each CLUDG operation provided there is an RPC method with matching annotations in the proto file. The action creator name will be the CLUDG verb followed by the state message field name (camelcase) followed by the async status (Request, Success, Failure, Cancel).

The *Request action creators will take the RPC input message type as a parameter. The *Success action creators will take the RPC output message type as a parameter. The *Failure actions will always pass the error string.

The *Request action will also have two additional optional parameters (resolve, reject). These callbacks will be called prior to dispatching the *Success or *Failure action from within the epic. This is useful for seeing if a specific request completed from within a component. 

*Request actions for verb "update" on a repeated field will take two parameters, prev and updated. Because protoc-gen-state has no knowledge of primary keys, the previous object is used to find the location of the object in the array within the reducer. 

For example, from proto files that look like this:
```
package package1

service BookService {
  rpc CreateBook(Book) returns (Book)
    rpc GetAllBooks(Book) returns (stream Book)
    rpc UpdateBook(Book) returns (Book)
    rpc DeleteBook(BookMetadata) returns (Book)
    rpc DoSomethingCrazy(package2.BirdInTheHand) returns (package2.OneInTheBush)
}

message Book { ... }
message BookMetadata { ... }
```

```
message ReduxState {
  option (state_options).type = REDUX_STATE;

  repeated package1.Book library = 1 [
    (state_field_options).method.create = "package1.CreateBook",
    (state_field_options).method.list = "package1.GetAllBook", // if repeated use list, otherwise use get
    (state_field_options).method.update = "package1.UpdateBook",
    (state_field_options).method.delete = "package1.DeleteBook"
  ];
}
message CustomActions {
  option (state_options).type = CUSTOM_ACTION;

  package1.Book doSomething = 1 [
    (state_field_options).method.custom = "package1.DoSomethingCrazy"
  ];
}
```

It will generate action creators for each of the annotations in the ReduxState message, assuming the RPC methods exist and are imported correctly (it will throw an error if not). Redux action creators do not return a value, but the call signatures are provided below.  

Notice the resolve/reject callbacks are optional.

  ```
  createLibraryRequest(
    library: package1.Book.AsObject, 
    resolve?: (payload: package1.Book.AsObject) => void, 
    reject?: (error: NodeJS.ErrnoException) => void
  );
  listLibraryRequest(
    library: package1.Book.AsObject,
    resolve?: (payload: package1.Book.AsObject[]) => void, 
    reject?: (error: NodeJS.ErrnoException) => void
  );
  updateLibraryRequest(
    prev: package1.Book.AsObject, 
    updated: package1.Book.AsObject,
    resolve?: (prev: package1.Book.AsObject, updated: package1.Book.AsObject) => void,
    reject?: (error: NodeJS.ErrnoException) => void,
  );
  deleteLibraryRequest(
    library: package1.BookMetadata.AsObject,
    resolve?: (payload: pacakge1.Book.AsObject) => void,
    reject?: (error: NodeJS.ErrnoException) => void,
  );

  createLibraryCancel()
  listLibraryCancel()
  updateLibraryCancel()
  deleteLibraryCancel()

  // also generates action creators for the async results of the api calls.
  // The epics dispatch these and the reducer catches them but they are not generally dispatched directly.
  // (Although they definitely could be if desired)
  createLibrarySuccess(library: package1.Book.AsObject)
  createLibraryFailure(error: NodeJS.ErrnoException)
  listLibrarySuccess(library: package1.Book.AsObject[]) // returns an array
  listLibraryFailure(error: NodeJS.ErrnoException)
  updateLibrarySuccess(library: {prev: package1.Book.AsObject, updated: package1.Book.AsObject}) // object with two keys
  updateLibraryFailure(error: NodeJS.ErrnoException)
  deleteLibrarySuccess(library: package1.Book.AsObject)
  deleteLibraryFailure(error: NodeJS.ErrnoException)

  // also generates one reset action per field
  // that sets the value in the reducer to the initial state
  resetLibrary()

  // and the custom actions (custom as in non-CRUD)
  // worth noting that a "custom" request without the callback parameters will not
  // give any feedback about whether it succeeded or not (no values in redux)
  // so these should almost always be passed with the callbacks
  customDoSomethingRequest(
    doSomething: package2.BirdInTheHand,
    resolve?: (payload: package2.OneInTheBush),
    reject?: (error: NodeJS.ErrnoException)
  );
  customDoSomethingRequestPromise() // deprecated, same as above
  customDoSomethingSuccess(doSomething: package2.OneInTheBush)
  customDoSomethingFailure(error: NodeJS.ErrnoException)
  customDoSomethingCancel()
  ```

  note: custom actions will not generate `reset` actions because they do not generate reducer cases
