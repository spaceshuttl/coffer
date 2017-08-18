import React from 'react'
import { Link } from 'react-router'

class Container extends React.Component {
  render() {
    return (
      <div>
        {this.props.children}
      </div>
    )
  }
}

export default Container
