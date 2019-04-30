import React, { Component } from 'react'
import loading from '../assets/loading.svg'
import { createMuiTheme } from '@material-ui/core/styles'

const theme = createMuiTheme({
  typography: {
    useNextVariants: true,
  },
  palette: {
    primary: {
      main: '#B432E6',
    },
    secondary: {
      main: '#3CBEAF',
    },
  },
})

function Callback() {
  return (
    <div>
      <img src={loading} alt='loading' />
    </div>
  )
}

export default Callback
