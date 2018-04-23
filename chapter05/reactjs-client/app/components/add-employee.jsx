import React, { Component, PropTypes } from 'react'

export default class AddEmployee extends React.Component {

  render(){
    return (
       <div>
          <input type = 'text' ref = 'input' />
          <button onClick = {(e) => this.handleClick(e)}>
             Add
          </button>
       </div>
    )

}
  handleClick(e) {
     const node = this.refs.input
     const text = node.value.trim()
     this.props.addEmployee(text)
     node.value = ''
  }
}
