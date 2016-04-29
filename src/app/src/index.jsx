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

    { /* Demo components */ }
    <row data-centered>
      <column cols="8">
        <div className="message message-warning" style={{ position: 'relative', maxWidth: 80 + 'vw' }}>
          Are you sure you want to <e>permanently</e> delete this? It'll be gone forever.
          <footer>
            <a href="#" className="btn" type="white" data-outline data-small> Yes!</a>
          </footer>
        </div>
      </column>
    </row>
  </div>,
  document.getElementById('app')
)
