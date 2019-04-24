import React, { useState } from 'react'
import Logo from 'components/Logo'
import IntegrationNotistack from './Snackbar'

import SoundSyncButton from 'components/Button'
import SoundSyncNavLink from 'components/NavLink'

import MusicControl from './MusicControl'
import {
  Button,
  IconButton,
  Typography,
  Toolbar,
  AppBar,
} from '@material-ui/core'
import MenuIcon from '@material-ui/icons/Menu'
import NotificationsIcon from '@material-ui/icons/Notifications'
import { withStyles } from '@material-ui/core/styles'

const styles = (theme) => {
  return {
    root: {
      flexGrow: 1,
    },
    grow: {
      flexGrow: 1,
    },
    container: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    textField: {
      marginLeft: theme.spacing.unit,
      marginRight: theme.spacing.unit,
      width: 200,
    },
    dense: {
      marginTop: 19,
    },
    menu: {
      width: 200,
    },
    title: {
      alignSelf: 'center',
    },
    logo: {
      alignSelf: 'flex-start',
    },
    code: {
      alignSelf: 'flex-end',
    },
    appBar: {
      top: 0,
    },
    musicControl: {
      bottom: 0,
      position: 'fixed',
    },
  }
}

function Party({ classes }) {
  const [values, setValues] = useState({
    isSpotifyLinked: false,
  })

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value })
  }

  const handleLinkSpotify = (e) => {
    e.preventDefault()
    console.log('Linking Spotify...')
  }

  return (
    <div className={classes.root}>
      <AppBar position='static' color='#F2F3F5'>
        <Toolbar>
          <Typography variant='h6' align='left'>
            Home
          </Typography>
          <Typography
            className={classes.grow}
            variant='h6'
            color='#262626'
            align='center'
            noWrap
          >
            soundsync
          </Typography>
          <Typography
            className={classes.grow}
            color='#262626'
            align='right'
            noWrap
          >
            <b>CODE: 4DR2</b>
          </Typography>
        </Toolbar>
      </AppBar>
      <AppBar position='static' color='#000000'>
        <Toolbar>
          <IconButton>
            <NotificationsIcon />
          </IconButton>
          <Typography
            className={classes.grow}
            variant='h6'
            color='inherit'
            noWrap
            align='center'
          >
            queue
          </Typography>
          <SoundSyncButton
            variant='contained'
            color='secondary'
            type='submit'
            className={classes.textField}
            onClick={handleLinkSpotify}
          >
            <SoundSyncNavLink
              color='inherit'
              to='/Party'
              className={classes.ButtonField}
            >
              LINK SPOTIFY
            </SoundSyncNavLink>
          </SoundSyncButton>
        </Toolbar>
      </AppBar>
      <IntegrationNotistack />
      <MusicControl className={classes.musicControl} />
    </div>
  )
}

export default withStyles(styles)(Party)
