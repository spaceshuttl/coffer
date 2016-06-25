import React, { PropTypes } from 'react'
import { Link } from 'react-router'

class Login extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      master: "",
    }
  }

  masterHandler(e) {
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
       master: this.state.master,
     }
   }
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
                <input type="password" className="input-big"></input>
              </section>
              <section>
                <Link to="manager">
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
