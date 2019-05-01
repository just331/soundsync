import React, { useState } from 'react'
import IntegrationNotistack from './Snackbar'

import SoundSyncButton from 'components/Button'
import SoundSyncNavLink from 'components/NavLink'

import Logo from '../assets/logo.png'
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
import Playlist from './Playlist'
import Search from './Search'

const styles = (theme) => {
  return {
    root: {
      flexGrow: 1,
      Typography: {
        fontFamily: ['Trebuchet MS', '"Helvetica"', 'Arial', 'sans-serif'],
      },
      button: {
        backgroundColor: 'rgba(255, 255, 255, 0.8)',
      },
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
      width: 20,
      paddingRight: 6,
    },
    code: {
      alignSelf: 'flex-end',
    },
    appBar: {
      top: 0,
      opacity: 0.5,
    },
    musicControl: {
      bottom: 0,
      position: 'fixed',
    },
    AppBar: {
      backgroundColor: 'rgba(255, 255, 255, 0.8)',
      color: 'black',
      position: 'static',
    },
  }
}

function Party({ classes }, props) {
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
      <AppBar className={classes.AppBar}>
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
            <img className={classes.logo} src={Logo} />
            soundsync
          </Typography>
          <Typography className={classes.grow} align='right' noWrap>
            <b>CODE: 4DR2</b>
          </Typography>
        </Toolbar>
      </AppBar>
      <AppBar className={classes.AppBar}>
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
      {/* <AppBar>
        <Search />
      </AppBar> */}
      <IntegrationNotistack />
      <Playlist />
      <MusicControl className={classes.musicControl} />
    </div>
  )
}

export default withStyles(styles)(Party)
