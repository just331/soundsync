import React from 'react'
import Grid from '@material-ui/core/Grid'
import { withStyles } from '@material-ui/core/styles'

const styles = (theme) => ({
  Container: {
    margin: 'auto',
    width: '100%',
    backgroundImage: `linear-gradient(${theme.palette.secondary.main}, ${
      theme.palette.primary.main
    })`,
    height: '100vh',
  },
})

function AppContainer({ children, classes }) {
  return (
    <Grid
      className={classes.Container}
      style={{ margin: 'auto', width: '100%' }}
      container
      spacing={24}
      children={children}
    />
  )
}

export default withStyles(styles)(AppContainer)
