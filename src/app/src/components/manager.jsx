'use strict'

import React, { PropTypes } from 'react'

var ws = new WebSocket("ws://localhost:5050");

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

    this.setState(function(previousState, currentProps) {
      return {accounts: previousState.accounts.push(this.state)};
    });
    console.log(this.state)
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
      var data = JSON.parse(event.data)
      for (var i = 0; i < data.length; i++) {
        data[i].key = data[i].identifier
      }
      console.log(data);
      this.setState({accounts: data})
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
  }

  deleteEntry(key) {
    console.log('deleting entry: ' + key)
    // TODO(mnzt): dispatch flux call to delete locally
  }

  render() {
    return (
      <tr>
        <td>{ this.props.identifier }</td>
        <td>{ this.props.password }
          <a onClick={ this.deleteEntry.bind(this, this.props.identifier) }  href="#" className="right">
            <i className="fa fa-times"/>
          </a>
        </td>
      </tr>
    )
  }
}

export { PasswordList, PasswordListAdd }
