import React, { useState, useReducer } from 'react'
import { withStyles } from '@material-ui/core/styles'

const initialState = {
  songs: [],
}

function reducer(state, action) {
  switch (action.type) {
    case 'add':
      return { songs: state.songs.push(newSong) }
  }
}

function PartyQueue({ initialState }) {
  const [state, dispatch] = useReducer(reducer, initialState)

  const handleAddSong = () => {
    //TODO: Listen for add song from SearchSong component

    // Add the song info to the state
    dispatch({ type: 'add' })
  }

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value })
  }

  return (
    <List>
      {values.songs.map((song) => {
        return <Song {...song} />
      })}
    </List>
  )
}

export default withStyles(styles)(PartyQueue)
