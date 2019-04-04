import React from 'react'
import Button from '@material-ui/core/Button'
import { withStyles } from '@material-ui/core/styles'

const styles = (theme) => {
  console.log(theme)
  return {
    Override: {
      borderRadius: 50,
      boxShadow: 'none',
    },
    Secondary: {
      backgroundColor: 'white',
    },
    ContainedSecondary: {
      color: 'white',
    },
  }
}

function SoundSyncButton({ children, classes, color, variant, onClick }) {
  return (
    <Button
      classes={{
        root: classes.Override,
        outlinedSecondary: classes.Secondary,
        containedSecondary: classes.ContainedSecondary,
      }}
      variant={variant}
      color={color}
      onClick={onClick}
    >
      {children}
    </Button>
  )
}

export default withStyles(styles)(SoundSyncButton)