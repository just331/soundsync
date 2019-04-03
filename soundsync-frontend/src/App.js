import React, { Component } from 'react';
import {
    Route,
    NavLink,
    HashRouter
} from "react-router-dom";
import logo from './logo.svg';
import CreateParty from './Create';
import JoinParty from "./JoinParty";
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
        <Button variant="contained" color="primary"><NavLink to="/CreateParty">Create Party</NavLink></Button>
        <Button><NavLink to="/">Join Party</NavLink></Button>
        <Route exact path="/" component={JoinParty}/>
        <Route path="/CreateParty" component={CreateParty}/>
        <Route path="/playlist" component={Playlist}/>
      </HashRouter>
    </MuiThemeProvider>
  );
}

export default App;
