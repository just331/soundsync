import React, { useState } from 'react'
import { withStyles } from '@material-ui/core/styles'

function Song({ classes }) {
  const [values, setValues] = useState({
    name: '',
    artist: '',
    songId: '',
    isPlaying: false,
  })

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value })
  }

  return (
    <ListItem>
      <ListItemAvatar>
        <Avatar>
          <ReorderIcon />
        </Avatar>
      </ListItemAvatar>
      <ListItemText primary={values.name} />
      <ListItemSecondaryAction>
        <IconButton>
          <DeleteIcon />
        </IconButton>
      </ListItemSecondaryAction>
    </ListItem>
  )
}

export default withStyles(styles)(Song)
