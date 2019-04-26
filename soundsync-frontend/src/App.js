import React from 'react'
import './index.css'
import { Route, HashRouter } from 'react-router-dom'
import CreateParty from 'components/CreateParty'
import JoinParty from 'components/JoinParty'
import Party from 'components/Party'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'

const theme = createMuiTheme({
  typography: {
    useNextVariants: true,
  },
  palette: {
    primary: {
      main: '#B432E6',
    },
    secondary: {
      main: '#3CBEAF',
    },
  },
})

function App() {
  return (
    <MuiThemeProvider theme={theme}>
      <HashRouter>
        <Route exact path='/' component={JoinParty} />
        <Route path='/CreateParty' component={CreateParty} />
        <Route path='/Party' component={Party} />
      </HashRouter>
    </MuiThemeProvider>
  )
}

export default App
