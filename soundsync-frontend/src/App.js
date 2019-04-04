import React from 'react'
import { Route, HashRouter } from 'react-router-dom'
import CreateParty from 'components/CreateParty'
import JoinParty from 'components/JoinParty'
import Party from 'components/Party'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import { Grid } from '@material-ui/core/'

const theme = createMuiTheme({
  typography: {
    useNextVariants: true,
  },
  spacing: {
    unit: 50,
  },
  textAlign: 'center',
})

function App() {
  return (
    <MuiThemeProvider theme={theme}>
      <Grid container justify='center' alignItems='center' spacing={16}>
        <HashRouter>
          <Route exact path='/' component={JoinParty} />
          <Route path='/CreateParty' component={CreateParty} />
          <Route path='/Party' component={Party} />
        </HashRouter>
      </Grid>
    </MuiThemeProvider>
  )
}

export default App
