import React from 'react';
import {
    Route,
    NavLink,
    HashRouter
} from "react-router-dom";
import CreateParty from './CreateParty';
import JoinParty from "./JoinParty";
import Party from "./Party";
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import {
  Grid,
  Button,
} from '@material-ui/core/';

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
          <Route exact path="/" component={JoinParty}/>
          <Route path="/CreateParty" component={CreateParty}/>
          <Route path="/Party" component={Party}/>
        </HashRouter>
      </Grid>
    </MuiThemeProvider>
  );
}

export default App;
