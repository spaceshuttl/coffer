'use strict'
import React, { PropTypes } from 'react'
import store from 'store'

// +1 to mrtbstyle
var ws = new WebSocket("ws://localhost:5050");

store.subscribe(() =>
  console.log(store.getState())
)

class PasswordListAdd extends React.Component {

  constructor(props) {
    super(props)
    this.state = props.initialCount
  }


  // On each key change, thes handleX functions will push the input field value to the store
  handleIdentifier(e) {
   this.setState({identifier: e.target.value});
  }

  handlePassword(e) {
     this.setState({password: e.target.value});
  }

  // handleSubmit inserts the data into store
  handleSubmit(e) {
    e.preventDefault()

    // set the action tag so the server know's what's up
    this.state.tag = "add"
    ws.send(JSON.stringify(this.state))

    console.log("dispathing " + this.state)
    state.dispatch({
      action: "ADD",
      account: this.state
    });

  }

  render() {
    return (
      <row data-centered>
        <column cols="8">
          <row data-centered data-end>
            <h4>Add Account</h4>
          </row>
          <row data-centered>
            <div className="btn-append">
              <input type="text"      name="key"    placeholder="Account/site"  onChange={ this.handleIdentifier.bind(this) }/>
              <input type="password"  name="value"  placeholder="Password"      onChange={ this.handlePassword.bind(this) }/>
              <span>
                <button class="btn" onClick={ this.handleSubmit.bind(this) }>Add</button>
              </span>
            </div>
          </row>
        </column>
      </row>
    )
  }
}

class PasswordList extends React.Component {

  constructor(props) {
    super(props)
    this.state = {
      accounts: []
    }

  }

  componentDidMount(){
    ws.onopen = () => ws.send(JSON.stringify({ tag: "all" }))
    ws.onmessage = (event) => {
      console.log(event.data);

      var data = JSON.parse(event.data)
      for (var i = 0; i < data.length; i++) {
        data[i].key = data[i].identifier
      }

      // HACK(mnzt): This is expensive and bad.
      data.map((account) => {
        state.dispatch({
          action: "ADD",
          account,
        })
      })
    }
  }

  render() {

    return (
    <row data-centered>
      <column cols="8">
        <table>
          <thead>
            <tr>
              <td class="width-6">Account</td>
              <td class="width-6">Password</td>
            </tr>
          </thead>
          <tbody>
            {this.state.accounts.map((account) => {
              return <PasswordListEntry key={ account.key } identifier={ account.identifier } password={ account.password }/>
            })}
          </tbody>
        </table>
      </column>
    </row>
  )
  }
}

class PasswordListEntry extends React.Component {
  constructor(props) {
    super(props);
    this.props = props
    this.state
  }

  deleteEntry(account) {
    var payload = {
      tag: "delete",
      identifier: account.identifier,
      password: null
    }

    ws.send(JSON.stringify(payload))

    state.dispatch({
      action: "DELETE",
      account,
    })

    console.log('deleting entry: ' + payload.identifier)
    // TODO(mnzt): dispatch flux call to delete locally
  }

  render() {
    return (
      <tr>
        <td>{ this.props.identifier }</td>
        <td>{ this.props.password }
          <a onClick={ this.deleteEntry.bind(this, this.props) }  href="#" className="right">
            <i className="fa fa-times"/>
          </a>
        </td>
      </tr>
    )
  }
}

export { PasswordList, PasswordListAdd }
