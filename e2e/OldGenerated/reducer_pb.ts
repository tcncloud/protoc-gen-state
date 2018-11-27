/* THIS FILE IS GENERATED FROM THE TOOL PROTOC-GEN-STATE  */
/* ANYTHING YOU EDIT WILL BE OVERWRITTEN IN FUTURE BUILDS */
import { getType, ActionType } from 'typesafe-actions';
import _ from 'lodash';
import * as protocActions from './actions_pb';
import * as ProtocTypes from './protoc_types_pb';
import { ProtocState, initialProtocState } from './state_pb';
type RootAction = ActionType<typeof protocActions>;


export function protocReducer(state: ProtocState = initialProtocState, action: RootAction) {
	switch(action.type) {
		case getType(protocActions['createLibraryRequest']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: true,
				}
			}
		case getType(protocActions['createLibrarySuccess']):
			var newLibraryValueArray: ProtocTypes.readinglist.Book.AsObject[] = [...state.library.value, action.payload] as ProtocTypes.readinglist.Book.AsObject[];
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
					value: newLibraryValueArray,
					error: initialProtocState.library.error,
				}
			}
		case getType(protocActions['createLibraryFailure']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
					error: { code: action.payload.code, message: action.payload.message },
				}
			}
		case getType(protocActions['createLibraryCancel']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
				}
			}
		case getType(protocActions['listLibraryRequest']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: true,
				}
			}
		case getType(protocActions['listLibrarySuccess']):
			var newLibraryValueArray: ProtocTypes.readinglist.Book.AsObject[] = action.payload;
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
					value: newLibraryValueArray,
					error: initialProtocState.library.error,
				}
			}
		case getType(protocActions['listLibraryFailure']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
					error: { code: action.payload.code, message: action.payload.message },
				}
			}
		case getType(protocActions['listLibraryCancel']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
				}
			}
		case getType(protocActions['updateLibraryRequest']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: true,
				}
			}
		case getType(protocActions['updateLibrarySuccess']):
			var newLibraryValueArray: ProtocTypes.readinglist.Book.AsObject[] = [...state.library.value] as ProtocTypes.readinglist.Book.AsObject[];
			var index: number = _.findIndex(newLibraryValueArray, action.payload.prev);
			if(index === -1){ newLibraryValueArray.push(action.payload.updated); } else {
				newLibraryValueArray[index] = action.payload.updated as ProtocTypes.readinglist.Book.AsObject;
			}
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
					value: newLibraryValueArray,
					error: initialProtocState.library.error,
				}
			}
		case getType(protocActions['updateLibraryFailure']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
					error: { code: action.payload.code, message: action.payload.message },
				}
			}
		case getType(protocActions['updateLibraryCancel']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
				}
			}
		case getType(protocActions['deleteLibraryRequest']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: true,
				}
			}
		case getType(protocActions['deleteLibrarySuccess']):
			var index: number = _.findIndex(state.library.value, action.payload);
			var newLibraryValueArray: ProtocTypes.readinglist.Book.AsObject[] = [...state.library.value.slice(0, index), ...state.library.value.slice(index+1)];
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
					value: newLibraryValueArray,
					error: initialProtocState.library.error,
				}
			}
		case getType(protocActions['deleteLibraryFailure']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
					error: { code: action.payload.code, message: action.payload.message },
				}
			}
		case getType(protocActions['deleteLibraryCancel']):
			return {
				...state,
				library: {
					...state.library,
					isLoading: false,
				}
			}
		case getType(protocActions['resetLibrary']):
			return {
				...state,
				library: initialProtocState.library
			}
		case getType(protocActions['createBookOfTheMonthRequest']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: true,
				}
			}
		case getType(protocActions['createBookOfTheMonthSuccess']):
			var newBookOfTheMonthValue: ProtocState["bookOfTheMonth"]["value"] = action.payload as ProtocState["bookOfTheMonth"]["value"];
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
					value: newBookOfTheMonthValue,
					error: initialProtocState.bookOfTheMonth.error,
				}
			}
		case getType(protocActions['createBookOfTheMonthFailure']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
					error: { code: action.payload.code, message: action.payload.message },
				}
			}
		case getType(protocActions['createBookOfTheMonthCancel']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
				}
			}
		case getType(protocActions['getBookOfTheMonthRequest']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: true,
				}
			}
		case getType(protocActions['getBookOfTheMonthSuccess']):
			var newBookOfTheMonthValue: ProtocState["bookOfTheMonth"]["value"] = action.payload;
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
					value: newBookOfTheMonthValue,
					error: initialProtocState.bookOfTheMonth.error,
				}
			}
		case getType(protocActions['getBookOfTheMonthFailure']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
					error: { code: action.payload.code, message: action.payload.message },
				}
			}
		case getType(protocActions['getBookOfTheMonthCancel']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
				}
			}
		case getType(protocActions['updateBookOfTheMonthRequest']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: true,
				}
			}
		case getType(protocActions['updateBookOfTheMonthSuccess']):
			var newBookOfTheMonthValue: ProtocState["bookOfTheMonth"]["value"] = { ...action.payload } as ProtocState["bookOfTheMonth"]["value"];
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
					value: newBookOfTheMonthValue,
					error: initialProtocState.bookOfTheMonth.error,
				}
			}
		case getType(protocActions['updateBookOfTheMonthFailure']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
					error: { code: action.payload.code, message: action.payload.message },
				}
			}
		case getType(protocActions['updateBookOfTheMonthCancel']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
				}
			}
		case getType(protocActions['deleteBookOfTheMonthRequest']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: true,
				}
			}
		case getType(protocActions['deleteBookOfTheMonthSuccess']):
			var newBookOfTheMonthValue: ProtocState["bookOfTheMonth"]["value"] = null;
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
					value: newBookOfTheMonthValue,
					error: initialProtocState.bookOfTheMonth.error,
				}
			}
		case getType(protocActions['deleteBookOfTheMonthFailure']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
					error: { code: action.payload.code, message: action.payload.message },
				}
			}
		case getType(protocActions['deleteBookOfTheMonthCancel']):
			return {
				...state,
				bookOfTheMonth: {
					...state.bookOfTheMonth,
					isLoading: false,
				}
			}
		case getType(protocActions['resetBookOfTheMonth']):
			return {
				...state,
				bookOfTheMonth: initialProtocState.bookOfTheMonth
			}
		case getType(protocActions['getTimeoutBookRequest']):
			return {
				...state,
				timeoutBook: {
					...state.timeoutBook,
					isLoading: true,
				}
			}
		case getType(protocActions['getTimeoutBookSuccess']):
			var newTimeoutBookValue: ProtocState["timeoutBook"]["value"] = action.payload;
			return {
				...state,
				timeoutBook: {
					...state.timeoutBook,
					isLoading: false,
					value: newTimeoutBookValue,
					error: initialProtocState.timeoutBook.error,
				}
			}
		case getType(protocActions['getTimeoutBookFailure']):
			return {
				...state,
				timeoutBook: {
					...state.timeoutBook,
					isLoading: false,
					error: { code: action.payload.code, message: action.payload.message },
				}
			}
		case getType(protocActions['getTimeoutBookCancel']):
			return {
				...state,
				timeoutBook: {
					...state.timeoutBook,
					isLoading: false,
				}
			}
		case getType(protocActions['resetTimeoutBook']):
			return {
				...state,
				timeoutBook: initialProtocState.timeoutBook
			}
		default: return state;
	}
};
