import React from 'react'
import './index.css'
import { Route } from 'react-router-dom'
import CreateParty from 'components/CreateParty'
import JoinParty from 'components/JoinParty'
import Party from 'components/Party'
import Callback from 'components/Callback'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import Auth from 'auth/Auth.js'

const auth = new Auth()

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

const handleAuthentication = (nextState, replace) => {
  if (/access_token|id_token|error/.test(nextState.location.hash)) {
    auth.handleAuthentication()
  }
}

class App extends React.Component {
  constructor(props) {
    super(props)
  }

  login() {
    this.props.auth.login()
  }

  logout() {
    this.props.auth.logout()
  }

  componentDidMount() {
    const { renewSession } = this.props.auth

    if (localStorage.getItem('isLoggedIn') === 'true') {
      renewSession()
    }
  }

  render() {
    return (
      <MuiThemeProvider theme={theme}>
        <Route
          exact
          path='/'
          render={(props) => <JoinParty auth={auth} {...props} />}
        />
        <Route
          path='/CreateParty'
          render={(props) => <CreateParty auth={auth} {...props} />}
        />
        <Route
          path='/Party/:partyId'
          render={(props) => <Party auth={auth} {...props} />}
        />
        <Route
          path='/Callback'
          render={(props) => {
            handleAuthentication(props)
            return <Callback {...props} />
          }}
        />
      </MuiThemeProvider>
    )
  }
}

export default App
