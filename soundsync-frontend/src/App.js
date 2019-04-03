import React, { Component } from 'react';
import {
    Route,
    NavLink,
    HashRouter
} from "react-router-dom";
import logo from './logo.svg';
import CreateParty from './Create';
import Home from "./Home";
import Playlist from "./Playlist";
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

const theme = createMuiTheme({
  spacing: {
    unit: 50,
  },
  textAlign: "center",
});

function App() {
  return (
    <MuiThemeProvider theme={theme}>
      <HashRouter>
        <Button variant="contained" color="primary"><NavLink to="/create">Create Party</NavLink></Button>
        <Button><NavLink to="/">Rejoin Party</NavLink></Button>
        <Route exact path="/" component={Home}/>
        <Route path="/create" component={CreateParty}/>
        <Route path="/playlist" component={Playlist}/>
      </HashRouter>
    </MuiThemeProvider>
  );
}

export default App;
