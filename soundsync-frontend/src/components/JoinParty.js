import React, { useState } from 'react'
import { Route, withRouter } from 'react-router-dom'
import Party from './Party'
import Grid from '@material-ui/core/Grid'
import SoundSyncButton from 'components/Button'
import SoundSyncNavLink from 'components/NavLink'
import SoundSyncInput from 'components/Input'
import AppContainer from 'components/AppContainer'
import { withStyles } from '@material-ui/core/styles'
import Logo from 'components/Logo'
import auth0Client from '../auth/Auth.js'

const styles = (theme) => {
  return {
    ButtonField: {
      width: 160,
      maxWidth: 500,
    },
    ButtonContainer: {
      display: 'flex',
      justifyContent: 'space-between',
      alignItems: 'center',
    },
    Input: {
      textAlign: 'center',
      padding: '0 !important',
    },
  }
}

function JoinParty({ classes }) {
  const [values, setValues] = useState({
    partyCode: '',
    nickName: '',
  })

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value })
  }

  // TODO: Make a post to join the party then redirect the user to the join party screen
  // TODO: Check to see if they have a valid jwt token first (means they are logged in)
  const handleJoinParty = (e) => {
    e.preventDefault()
  }

  return (
    <AppContainer>
      <Grid className={classes.ButtonContainer} item xs={12}>
        <SoundSyncButton color='secondary' variant='outlined'>
          <SoundSyncNavLink color='secondary' to='/'>
            Rejoin As Host
          </SoundSyncNavLink>
        </SoundSyncButton>
        <SoundSyncButton color='primary' variant='contained'>
          <SoundSyncNavLink color='inherit' to='/CreateParty'>
            Create Party
          </SoundSyncNavLink>
        </SoundSyncButton>
      </Grid>
      <Logo className={classes.Logo} />
      <Grid className={classes.Input} item xs={12}>
        <SoundSyncInput
          id='partyCode'
          placeholder='Party Code'
          value={values.partyCode}
          onChange={handleChange('partyCode')}
        />
        <SoundSyncInput
          id='nickName'
          placeholder='Nickname'
          value={values.nickName}
          onChange={handleChange('nickName')}
        />
        <SoundSyncInput
          id='phoneNum'
          placeholder='Phone Number'
          value={values.phoneNum}
          onChange={handleChange('phoneNum')}
        />
        <Route path='/Party/4DR2' component={Party} />
        <SoundSyncButton
          variant='contained'
          color='secondary'
          type='submit'
          className={classes.textField}
          onClick={handleJoinParty}
        >
          <SoundSyncNavLink
            color='inherit'
            to='/Party/4DR2'
            className={classes.ButtonField}
          >
            Join Party
          </SoundSyncNavLink>
        </SoundSyncButton>
      </Grid>
    </AppContainer>
  )
}

export default withRouter(withStyles(styles)(JoinParty))
