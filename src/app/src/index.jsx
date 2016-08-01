import ReactDOM from 'react-dom'
import { PasswordList, PasswordListAdd } from './components/manager'


ReactDOM.render(
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
  </div>,
  document.getElementById('app')
)
