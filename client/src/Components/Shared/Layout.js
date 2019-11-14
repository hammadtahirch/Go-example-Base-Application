import React, { Component } from 'react'
import Header from './Header.js'
import Footer from './Footer.js'
import Login from '../Pages/Account/Login.js'
class Layout extends Component {
  render() {
    console.log(this.props)
    return (
      <div>
        <Header />
        {this.props.children}
        <Footer/>
      </div>
    )
  }
}
export default Layout