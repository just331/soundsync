import React, { useState } from 'react'
import { Route } from 'react-router-dom'
import Party from './Party'
import TextField from '@material-ui/core/TextField'
import Grid from '@material-ui/core/Grid'
import SoundSyncButton from 'components/Button'
import SoundSyncNavLink from 'components/NavLink'
import { withStyles } from '@material-ui/core/styles'

const styles = (theme) => {
  return {
    textField: {
      width: '86vw',
    },
    ButtonContainer: {
      display: 'flex',
      justifyContent: 'space-between',
      alignItems: 'center',
    },
    Input: {
      textAlign: 'center',
    },
    Container: {
      margin: 'auto',
      width: '100%',
      backgroundImage: `linear-gradient(${theme.palette.secondary.main}, ${
        theme.palette.primary.main
      })`,
      padding: '160px 0px',
      height: '100vh',
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
    console.log('Joined Party!')
  }

  return (
    <Grid
      className={classes.Container}
      style={{ margin: 'auto', width: '100%' }}
      container
      spacing={24}
    >
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
      <Grid className={classes.Input} item xs={12}>
        <TextField
          id='partyCode'
          label='Party Code'
          required
          className={classes.textField}
          value={values.partyCode}
          onChange={handleChange('partyCode')}
        />
      </Grid>
      <Grid className={classes.Input} item xs={12}>
        <TextField
          id='nickName'
          label='Nickname'
          required
          className={classes.textField}
          value={values.nickName}
          onChange={handleChange('nickName')}
        />
      </Grid>
      <Grid className={classes.Input} item xs={12}>
        <Route path='/Party' component={Party} />
        <SoundSyncButton
          variant='contained'
          color='secondary'
          type='submit'
          className={classes.textField}
          onClick={handleJoinParty}
        >
          <SoundSyncNavLink color='inherit' to='/Party'>
            Join Party
          </SoundSyncNavLink>
        </SoundSyncButton>
      </Grid>
    </Grid>
  )
}

export default withStyles(styles)(JoinParty)
