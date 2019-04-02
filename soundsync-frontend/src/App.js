import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import CreateParty from './Create';
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';

const theme = createMuiTheme({
  spacing: {
    unit: 50,
  },
  text-align: center,
});

function App() {
  return (
    <MuiThemeProvider theme={theme}>
      <h1>soundsync</h1>
      <CreateParty/>
    </MuiThemeProvider>
  );
}

export default App;
