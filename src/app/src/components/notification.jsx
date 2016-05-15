'use strict'
import React, { PropTypes } from 'react'

class Notification extends React.Component {
  constructor(props) {
    super(props)
  }

  render(){
    return (
      <div className={"message message-" + this.props.level}>{this.props.message}</div>
    )
  }
}

export default Notification
