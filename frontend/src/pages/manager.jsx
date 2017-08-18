import React from 'react'
import Notification from '../components/notification'
import { PasswordList, PasswordListAdd } from '../components/manager'

class Manager extends React.Component {
  render () {
    return (
      <div>
        { /* Greeting */ }
        <row data-centered style={{ marginTop: 6 + 'rem' }}>
          <column cols="8" className="text-centered">
            <h1>Coffer</h1>
            <p>Your light-weight, easy to use, secure, password manager</p>
          </column>
        </row>

        <PasswordListAdd />
        <PasswordList />
      </div>
    )
  }
}

export default Manager
