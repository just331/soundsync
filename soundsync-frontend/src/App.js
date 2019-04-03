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
import Grid from '@material-ui/core/Grid';
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
      <Grid container justify="center" alignItems="center" spacing={16}>
        <HashRouter>
          <Button variant="contained" color="primary"><NavLink to="/CreateParty">Create Party</NavLink></Button>
          <Button variant="contained" color="primary"><NavLink to="/">Join Party</NavLink></Button>
          <Route exact path="/" component={JoinParty}/>
          <Route path="/CreateParty" component={CreateParty}/>
          <Route path="/playlist" component={Playlist}/>
        </HashRouter>
      </Grid>
    </MuiThemeProvider>
  );
}

export default App;
