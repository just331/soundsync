import React, { useState } from 'react'
import { withStyles } from '@material-ui/core/styles'

const styles = (theme) => {
  return {}
}

function queueSong(song) {
  //create a unike key for each new fruit item
  var timestamp = new Date().getTime()
  // update the state object
  this.state.songs['song-' + timestamp] = song
  // set the state
  this.setState({ songs: this.state.songs })
}

function Playlist({ classes }, props) {
  const [values, setValues] = useState({
    songs: { hello: 'world' },
  })

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value })
  }

  return (
    <div className='container'>
      <ul className='list-group text-center'>
        {Object.keys(values.songs).map(
          function(key) {
            return (
              <li className='list-group-item list-group-item-info'>
                {values.songs[key]}
              </li>
            )
          }.bind(this),
        )}
      </ul>
    </div>
  )
}

export default withStyles(styles)(Playlist)
