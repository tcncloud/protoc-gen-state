
import e2e_redux3_protos_readinglist_readinglist from '../readinglist/readinglist_pb';
import google_protobuf_timestamp from 'google-protobuf/google/protobuf/timestamp_pb';
const messageMap = new Map();

const readinglist_Book_map = new Map();
readinglist_Book_map.set('creationDate', google_protobuf_timestamp.Timestamp);
messageMap.set(e2e_redux3_protos_readinglist_readinglist.Book, readinglist_Book_map);


function getNestedMessageConstructor(messageType, fieldName) {
  return messageMap.has(messageType) && messageMap.get(messageType).get(fieldName);
}

export function toMessage(obj: any, messageClass: any) {
  if (!obj) {
    return new messageClass();
  }
  
  const message = new messageClass();

  Object.keys(obj).forEach(key => {
    
    let ele = obj[key];
    const upperCaseKey = key.charAt(0).toUpperCase() + key.substr(1);
    const setterName = "set" + upperCaseKey;
    const getterName = "get" + upperCaseKey;
    

    if (message[setterName]) {
      var nestedMessageContructor = getNestedMessageConstructor(messageClass, key);
      if (nestedMessageContructor) {
        if (key.length > 4 && key.slice(key.length - 4) === 'List' && Array.isArray(ele)) { // check if field is repeated
          
          ele = ele.map(subEle => toMessage(subEle, nestedMessageContructor));
        } else {
          
          ele = toMessage(ele, nestedMessageContructor);
        }
      }

      message[setterName](ele);
    } else if (message[getterName] && key.slice(key.length - 3) === 'Map') { // check if field is a map
      
      // if the map field is missing, nothing needs to be done.
      if (ele !== undefined && ele !== null) {
        if (Array.isArray(ele)) {
          if (ele.length) {
            var mapObj = message[getterName]();
            var mappedFieldValueConstructor = getNestedMessageConstructor(messageClass, key);
            if (mappedFieldValueConstructor) {
              
              ele = ele.map(([key, value]) => {
                
                return [key, mappedFieldValueConstructor(value)];
              });
              
            }
            
            ele.forEach(([key, value]) => {
              
              mapObj.set(key, value);
            });
            
          }
        } else {
          
          throw new Error("Protoc-gen-state: Expected field " + key + " to be an array of tuples.");
        }
      }
    } else {
      
      throw new Error("No corresponding gRPC setter method for given key: " + key);
    }
    
  });
  

  return message;
}
