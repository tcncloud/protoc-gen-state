import * as React from "react";
import { connect } from "react-redux";
import { Dispatch, bindActionCreators  } from 'redux';
import _ from 'lodash';

import { RootAction } from '../rootAction';
import { RootState } from '../rootState';
import { ProtocState } from 'protos/BasicState/state_pb';
import * as ProtocActions from 'protos/BasicState/actions_pb';
import * as ProtocTypes from 'protos/BasicState/protoc_types_pb';


console.log("PROTOCTYPES: ", ProtocTypes)
console.log("PROTOCACTIONS: ", ProtocActions)

  interface PropsFromState {
  protocLibrary: ProtocState["library"];
  bookOfTheMonth: ProtocState["bookOfTheMonth"];
};

interface PropsFromDispatch {
  // complex type
  protocCreateBookOfTheMonthRequest: (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => void;
  protocGetBookOfTheMonthRequest: (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => void;
  protocUpdateBookOfTheMonthRequest: (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => void;
  protocDeleteBookOfTheMonthRequest: (bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject) => void;

  // complex array
  protocCreateLibraryCancel: () => void;
  protocCreateLibraryRequest: (library: ProtocTypes.readinglist.Book.AsObject) => void;
  protocListLibraryRequest: (library: ProtocTypes.readinglist.Book.AsObject) => void;
  protocUpdateLibraryRequest: (prev: ProtocTypes.readinglist.Book.AsObject, updated: ProtocTypes.readinglist.Book.AsObject) => void;
  protocDeleteLibraryRequest: (library: ProtocTypes.readinglist.Book.AsObject) => void;
}

interface ComponentLocalState {
  addBookBox: ProtocTypes.readinglist.Book.AsObject,
  bookOfTheMonth: ProtocTypes.readinglist.Book.AsObject,
}

interface ReduxProps extends PropsFromState, PropsFromDispatch {}


class BookComponent extends React.Component<ReduxProps, ComponentLocalState> {
  state: ComponentLocalState;
  constructor(props: ReduxProps) {
    super(props);

    this.state = {
      addBookBox: { title: "Ready Player One", author: "Ernest Cline" } as ProtocTypes.readinglist.Book.AsObject,
      bookOfTheMonth: { title: "Ulysses", author: "James Joyce" } as ProtocTypes.readinglist.Book.AsObject,
    };
  }

  updateBookOfTheMonthTitle = (myVal : string) => {
    this.setState({
      bookOfTheMonth: {
        ...this.state.bookOfTheMonth,
        title: myVal
      }
    });
  }
  updateBookOfTheMonthAuthor = (myVal : string) => {
    this.setState({
      bookOfTheMonth: {
        ...this.state.bookOfTheMonth,
        author: myVal
      }
    });
  }
  updateAddBookBoxTitle = (myVal : string) => {
    this.setState({
      addBookBox: {
        ...this.state.addBookBox,
        title: myVal
      }
    });
  }
  updateAddBookBoxAuthor = (myVal: string) => {
    this.setState({
      addBookBox: {
        ...this.state.addBookBox,
        author: myVal
      }
    });
  }

  componentWillReceiveProps() {
  }

  render() {
    console.log('idk', ProtocTypes);
    return (
      <div style={{background: 'white'}} className="App">
        <div>
        <h1> Complex Type </h1>
        <div>
          <label>
            Title
            <input style={{"margin":"10px"}} type="text" value={this.state.bookOfTheMonth.title} onChange={(e) => this.updateBookOfTheMonthTitle(e.target.value)} name="addBookTitle" />
          </label>
          <label>
            Author
            <input style={{"margin":"10px"}} type="text" value={this.state.bookOfTheMonth.author} onChange={(e) => this.updateBookOfTheMonthAuthor(e.target.value)} name="addBookAuthor" />
          </label>

          <br />
          {this.props.bookOfTheMonth && this.props.bookOfTheMonth.value && <p> "{this.props.bookOfTheMonth.value.title}" by {this.props.bookOfTheMonth.value.author} </p>}
          <br />

        </div>
        <button onClick={() => this.props.protocCreateBookOfTheMonthRequest(this.state.bookOfTheMonth)}>Create</button>
        <button onClick={() => this.props.protocGetBookOfTheMonthRequest(this.state.bookOfTheMonth)}>Read</button>
        <button onClick={() => this.props.protocUpdateBookOfTheMonthRequest(this.state.bookOfTheMonth)}>Update</button>
        <button onClick={() => this.props.protocDeleteBookOfTheMonthRequest(this.state.bookOfTheMonth)}>Delete</button>
        { this.props.bookOfTheMonth && this.props.bookOfTheMonth.error != null && <p> Error: {this.props.bookOfTheMonth.error.message}</p>}
      </div>

      <div>
        <h1> Complex Type Array</h1>
        <form onSubmit={(e) => e.preventDefault()}>

          <label style={{"padding":"10px"}}>
            Title
            <input style={{"margin":"10px"}} type="text" value={this.state.addBookBox.title} onChange={(e) => this.updateAddBookBoxTitle(e.target.value)} name="add-book" />
          </label>
          <label style={{"padding":"10px"}}>
            Author
            <input style={{"margin":"10px"}} type="text" value={this.state.addBookBox.author} onChange={(e) => this.updateAddBookBoxAuthor(e.target.value)} name="add-book" />
          </label>

          <br />
          <button onClick={() => this.props.protocCreateLibraryRequest(this.state.addBookBox)}>Create</button>
          <button onClick={() => console.log('props: ', this.props)}>Props</button>
          <button onClick={() => this.props.protocCreateLibraryCancel()}>Cancel</button>
          <button onClick={() => this.props.protocListLibraryRequest(this.state.bookOfTheMonth)}>Read</button>
        </form>

        {this.props.protocLibrary.isLoading && <button disabled>fetching...</button>}
        {this.props.protocLibrary.value && _.map(this.props.protocLibrary.value, (book, i) => {
        return (
        <div key={book.title + book.author}>{book.title} by {book.author}
          <button key={book.title + book.author + "update"}
            onClick={() => this.props.protocUpdateLibraryRequest(this.props.protocLibrary.value[i], this.state.addBookBox)}
          >Update</button>
          <button key={book.title + book.author + "delete"}
          onClick={() => {
            this.props.protocDeleteLibraryRequest(this.props.protocLibrary.value[i])
          }}
          >Delete</button>
        </div>
      )
        })}
      </div>
      </div>
    );
  }
}

function mapStateToProps(state: RootState): PropsFromState {
  return {
    protocLibrary: state.protoc.library,
    bookOfTheMonth: state.protoc.bookOfTheMonth,
  };
};

function mapDispatchToProps(dispatch: Dispatch<RootAction>): PropsFromDispatch {
  return bindActionCreators({
    // complex type
    protocCreateBookOfTheMonthRequest: ProtocActions.createBookOfTheMonthRequest,
    protocGetBookOfTheMonthRequest: ProtocActions.getBookOfTheMonthRequest,
    protocUpdateBookOfTheMonthRequest: ProtocActions.updateBookOfTheMonthRequest,
    protocDeleteBookOfTheMonthRequest: ProtocActions.deleteBookOfTheMonthRequest,

    // complex array
    protocCreateLibraryCancel: ProtocActions.createLibraryCancel,
    protocCreateLibraryRequest: ProtocActions.createLibraryRequest,
    protocListLibraryRequest: ProtocActions.listLibraryRequest,
    protocUpdateLibraryRequest: ProtocActions.updateLibraryRequest,
    protocDeleteLibraryRequest: ProtocActions.deleteLibraryRequest,
  }, dispatch);
}

export default connect<PropsFromState, PropsFromDispatch, {}>(mapStateToProps, mapDispatchToProps)(BookComponent);
