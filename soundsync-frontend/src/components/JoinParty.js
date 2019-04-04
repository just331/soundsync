import React, { useState } from 'react'
import { Route } from 'react-router-dom'
import Party from './Party'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import SoundSyncButton from 'components/Button'
import SoundSyncNavLink from 'components/NavLink'
import { withStyles } from '@material-ui/core/styles'
import Logo from 'assets/logo.png'

const styles = (theme) => {
  return {
    LogoText: {
      color: 'white',
    },
    ImageContainer: {
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
      flexDirection: 'column',
    },
    Image: {
      width: '45vw',
    },
    ButtonField: {
      width: '300px',
    },
    textField: {
      placeholder: 'name',
      color: '#333333',
      height: '1vh',
      backgroundColor: 'white',
      borderRadius: '50px',
      boxShadow: 'none',
      border: '2px solid white',
      padding: '20px',
      margin: '10px 0px',
      width: '300px',
      '&::placeholder': {
        color: '#666666',
      },
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
      <Grid className={classes.ImageContainer} item xs={12}>
        <img className={classes.Image} src={Logo} alt='logo' />
        <Typography className={classes.LogoText} variant='h4' gutterBottom>
          soundsync
        </Typography>
      </Grid>
      <Grid className={classes.Input} item xs={12}>
        <input
          id='partyCode'
          placeholder='Party Code'
          className={classes.textField}
          value={values.partyCode}
          onChange={handleChange('partyCode')}
        />
        <input
          id='nickName'
          placeholder='Nickname'
          className={classes.textField}
          value={values.nickName}
          onChange={handleChange('nickName')}
        />
        <Route path='/Party' component={Party} />
        <SoundSyncButton
          variant='contained'
          color='secondary'
          type='submit'
          className={classes.textField}
          onClick={handleJoinParty}
        >
          <SoundSyncNavLink
            color='inherit'
            to='/Party'
            className={classes.ButtonField}
          >
            Join Party
          </SoundSyncNavLink>
        </SoundSyncButton>
      </Grid>
    </Grid>
  )
}

export default withStyles(styles)(JoinParty)
