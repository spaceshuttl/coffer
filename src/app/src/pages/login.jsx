import React, { PropTypes } from 'react'
import { Link } from 'react-router'
import { ws } from '../components/websocket.jsx'

class Login extends React.Component {
  constructor(props) {
    super(props)
    this.props = props

    this.state = {
      master: "",
    }
  }

  handleMaster(e) {
    e.preventDefault()
    this.setState({master: e.target.value});
  }

  handleKeypress(e) {
   if (e.key == 'Enter') {
     this.handleSubmit(e)
   }
  }

  // handleSubmit inserts the data into store
  handleSubmit(e) {
   e.preventDefault()

   let request = {
     action: "LOGIN",
     payload: {
       password: this.state.master,
     }
   }

   console.log(request);
   ws.send(JSON.stringify(request))

   // on success coninue to the manager
   this.props.history.push('manager');
  }

  render () {
    return (
      <div>
        <row data-centered>
            <column cols="12" className="text-centered">
              <h2>Login</h2>
            </column>
        </row>
        <row data-centered>
          <column cols="8" className="text-centered">
            <form className="forms">
              <section>
                <input
                  type="password"
                  className="input-big"
                  onChange={ this.handleMaster.bind(this) }
                  onKeyPress={ this.handleKeypress.bind(this) }
                ></input>
              </section>
              <section>
                <Link to="manager" onClick={this.handleSubmit.bind(this)}>
                  <button type="primary">Log in</button>
                </Link>
              </section>
            </form>
          </column>
        </row>
      </div>
    )
  }

}

export default Login
