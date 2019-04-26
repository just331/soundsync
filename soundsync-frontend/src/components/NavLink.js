import React from 'react'
import { Link as RouterLink } from 'react-router-dom'
import Link from '@material-ui/core/Link'
import { withStyles } from '@material-ui/core/styles'

const styles = (theme) => ({})

function SoundSyncNavLink(props) {
  return (
    <Link
      color={props.color}
      component={RouterLink}
      underline='none'
      {...props}
    />
  )
}

export default withStyles(styles)(SoundSyncNavLink)
