/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */

import * as ProtocTypes from './protoc_types_pb';

export interface ProtocState { 
library: {
  isLoading: boolean;
  error: { code: string; message: string; };
  value: ProtocTypes.readinglist.Book.AsObject[];
  
},

bookOfTheMonth: {
  isLoading: boolean;
  error: { code: string; message: string; };
  value: ProtocTypes.readinglist.Book.AsObject | null;
},

timeoutBook: {
  isLoading: boolean;
  error: { code: string; message: string; };
  value: ProtocTypes.readinglist.Book.AsObject | null;
},

}

export const initialProtocState : ProtocState = { 
library: {
  isLoading: false,
  error: null,
  value: [],
  
},

bookOfTheMonth: {
  isLoading: false,
  error: null,
  value: null,
},

timeoutBook: {
  isLoading: false,
  error: null,
  value: null,
},

}
