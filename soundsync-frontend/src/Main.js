import React, { Component } from "react";
import {
    Route,
    NavLink,
    HashRouter
} from "react-router-dom";
import Home from "./Home";
import Create from "./Create";
import Playlist from "./Playlist";
import Button from '@material-ui/core/Button';

class Main extends Component {
  render() {
    return (
        <HashRouter>
          <div className="navbar">
            <Button><NavLink to="/create">Create Party</NavLink></Button>
            <Button><NavLink to="/">Rejoin Party</NavLink></Button>
          </div>
          <div className="content">
            <Route exact path="/" component={Home}/>
            <Route path="/create" component={Create}/>
            <Route path="/playlist" component={Playlist}/>
          </div>
        </HashRouter>
    );
  }
}

export default Main;
