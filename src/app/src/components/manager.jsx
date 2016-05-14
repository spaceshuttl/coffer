'use strict'
import React, { PropTypes } from 'react'

// +1 to mrtbstyle
var ws = new WebSocket("ws://localhost:5050");

class PasswordListAdd extends React.Component {

  constructor(props) {
    super(props)
    this.state = {
      key: "",
      identifier: "",
      value: ""
    }
  }


  // On each key change, thes handleX functions will push the input field value to the store
  handleIdentifier(e) {
   this.setState({identifier: e.target.value});
  }

  handlePassword(e) {
     this.setState({password: e.target.value});
  }

  handleKeypress(e) {
    if (e.key == 'Enter') {
      this.handleSubmit(e)
    }
  }

  // handleSubmit inserts the data into store
  handleSubmit(e) {
    e.preventDefault()

    var key = Math.random().toString(36).substring(24);

    let request = {
      action: "ADD",
      payload: {
        key,
        identifier: this.state.identifier,
        value: this.state.password
      }
    }

    ws.send(JSON.stringify(request))

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
              <input  ref="identifier"
                      type="text"
                      name="key"
                      placeholder="Account/site"
                      onChange={ this.handleIdentifier.bind(this) }
              />
              <input  ref="password"
                      type="password"
                      name="value"
                      placeholder="Password"
                      onChange={ this.handlePassword.bind(this) }
                      onKeyPress={ this.handleKeypress.bind(this) }
              />
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

    let request = {
      action: "ALL"
    }

    ws.onopen = () => {
      ws.send(JSON.stringify(request))
    }
  }

  componentDidMount() {
    ws.addEventListener('message', (event) => {
      let response = JSON.parse(event.data)

      // TODO(mnzt): implement error handling
      // if (response.error) {
      //   console.error(data)
      // }
      this.setState({
        accounts: response
      })

    })
  }

  render() {

    let accounts = this.state.accounts.map((account) => {
      return <PasswordListEntry key={ account.key } _key={ account.key } identifier={ account.identifier } password={ account.value }/>
    })

    return (
    <row data-centered>
      <column cols="8">
        <table>
          <thead>
            <tr>
              <td className="width-4 text-centered">Account</td>
              <td className="width-5 text-centered">Password</td>
              <td className="width-3"></td>
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
    this.state = {
      display: { WebkitTextSecurity: 'disc' }
    }
  }

  deleteEntry(account) {
    let data = {
      action: "DELETE",
      payload: {
        key: account._key,
      }
    }

    console.log(data)
    ws.send(JSON.stringify(data))
  }

  showPassword(event) {
    if (JSON.stringify(this.state.display) === JSON.stringify({WebkitTextSecurity: "none"})) {
      this.setState({display: { WebkitTextSecurity: "disc" }})
      return
    }
    this.setState({display: { WebkitTextSecurity: "none" }})
  }

  render() {
    return (
      <tr>
        <td>
          { this.props.identifier }
        </td>
        <td ref={this.props._key} style={ this.state.display }>
          { this.props.password }
        </td>
        <td>
          <span className="btn-group right">
            <button data-small data-outline onClick={ this.showPassword.bind(this) }>
              <i className="fa fa-eye"/>
            </button>
            <button type="success" id="cpy" data-clipboard-action="copy" data-clipboard-text={ this.props.password} data-small data-outline >
              <i className="fa fa-clipboard" />
            </button>
            <button className="btn" type="red" data-small data-outline >
              <i className="fa fa-trash" />
            </button>
          </span>
        </td>
      </tr>
    )
  }
}

export { PasswordList, PasswordListAdd }
