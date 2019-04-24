import React from 'react'
import SSLogo from 'assets/logo.png'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import { withStyles } from '@material-ui/core/styles'

const styles = (theme) => ({
  LogoText: {
    color: 'white',
    fontSize: '5.6vw',
    marginBottom: 0,
    typography: {
      fontFamily: [
        '-apple-system',
        'BlinkMacSystemFont',
        '"Segoe UI"',
        'Roboto',
        '"Helvetica"',
        'Arial',
        'sans-serif',
        '"Apple Color Emoji"',
        '"Segoe UI Emoji"',
        '"Segoe UI Symbol"',
      ].join(','),
    }
  },
  ImageContainer: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    flexDirection: 'column',
    padding: '0 !important',
  },
  Image: {
    width: '65%',
    maxWidth: 500,
  },
})

function Logo({ classes }) {
  return (
    <Grid className={classes.ImageContainer} item xs={12}>
      <img className={classes.Image} src={SSLogo} alt='logo' />
      <Typography className={classes.LogoText} variant='h4' gutterBottom>
        soundsync
      </Typography>
    </Grid>
  )
}

export default withStyles(styles)(Logo)
