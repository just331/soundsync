import React, { Component } from "react";
import Logo from './logo.png';
import {
    Route,
    NavLink,
    HashRouter
} from "react-router-dom";
import Stuff from "./Create";
import Playlist from "./Playlist";
 
class Home extends Component {
    constructor() {
        super();
        this.state = {
            customers: [],
            open: false,
            explicitswitch: true,
            style : {width: 0}
    };
    }
  
    render() {
        return (
        <div>
        <img src={Logo} alt="soundsync Logo" class="logo"></img>
            <h1>soundsync</h1>
            <form method="post">
                <input
                name="Name"
                type="string"
                placeholder="Enter Party Code"
                required="required"
                value={this.state.numberOfGuests}
                onChange={this.handleInputChange} />
                <br></br>
                <input
                name="Name"
                type="string"
                placeholder="Enter Username"
                required="required"
                value={this.state.numberOfGuests}
                onChange={this.handleInputChange} />
                <br></br>
                <NavLink class="joined" to="/playlist">Join Party</NavLink>
            </form>
      </div>
    );
  }
}
 
export default Home;