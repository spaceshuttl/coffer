'use strict'
import React, { PropTypes } from 'react'

// +1 to mrtbstyle
var ws = new WebSocket("ws://localhost:5050");

class PasswordListAdd extends React.Component {

  constructor(props) {
    super(props)
    this.state = {
      identifier: "",
      password: ""
    }
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

    console.log(e);

    // set the action tag so the server know's what's up
    var payload = {
      tag: "ADD",
      identifier: this.state.identifier,
      password: this.state.password
    }

    ws.send(JSON.stringify(payload))

   // reset the form
   ReactDOM.findDOMNode(this.refs.identifier).value = ""; // Unset the value
   ReactDOM.findDOMNode(this.refs.password).value = ""; // Unset the value

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
              <input ref="identifier" type="text"      name="key"    placeholder="Account/site"  onChange={ this.handleIdentifier.bind(this) }/>
              <input ref="password" type="password"  name="value"  placeholder="Password"      onChange={ this.handlePassword.bind(this) }/>
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

  constructor() {
    super()

    this.state = {
      accounts: []
    }

    ws.onopen = () => {
      ws.send(JSON.stringify({
        tag: "ALL"
      }))
    }
  }

  componentDidMount() {
    ws.addEventListener('message', (event) => {
      var response = JSON.parse(event.data)

      if (response.error) {
        console.error(data)
      }

      var data = response.data

      data.map((account) => {
        account.key = account.identifier
      })

      this.setState({
        accounts: data
      })

    })
  }

  render() {

    let accounts = this.state.accounts.map((account) => {
      return <PasswordListEntry key={ account.key } identifier={ account.identifier } password={ account.password }/>
    })

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
            { accounts }
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
      tag: "DEL",
      identifier: account.identifier,
      password: null
    }

    ws.send(JSON.stringify(payload))

    console.log('deleting entry: ' + payload.identifier)
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
