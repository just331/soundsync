import React, { Component } from "react";
import {
    Route,
    NavLink,
    HashRouter
} from "react-router-dom";
import Home from "./Home";
import Create from "./Create";
import Playlist from "./Playlist";
 
class Main extends Component {
  render() {
    return (
        <HashRouter>
        <div>
          <div className="navbar">
            <ul>
            <li><NavLink to="/create">Create Party</NavLink></li>
            <li><NavLink to="/">Rejoin Party</NavLink></li>
            </ul>
          </div>
          <div className="content">
            <Route exact path="/" component={Home}/>
            <Route path="/create" component={Create}/>
            <Route path="/playlist" component={Playlist}/>
          </div>
        </div>
        </HashRouter>
    );
  }
}
 
export default Main;