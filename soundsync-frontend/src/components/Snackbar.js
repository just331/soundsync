import React from 'react'
import Snackbar from '@material-ui/core/Button'
import { withStyles } from '@material-ui/core/styles'
import Button from '@material-ui/core/Button'
import PropTypes from 'prop-types'
import { SnackbarProvider, withSnackbar } from 'notistack'
import Slide from '@material-ui/core/Slide'

const styles = (theme) => ({
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
})

function TransitionUp(props) {
  return <Slide {...props} direction='up' />
}

function TransitionDown(props) {
  return <Slide {...props} direction='down' />
}

class App extends React.Component {
  state = {
    open: false,
    Transition: null,
  }

  handleClick = (Transition) => () => {
    this.props.enqueueSnackbar('XYZ song was added to playlist.')
  }

  handleClose = () => {
    this.setState({ open: false })
  }

  render() {
    return (
      <React.Fragment>
        <Button onClick={this.handleClick(TransitionUp)}>
          Click to add XYZ song.
        </Button>
      </React.Fragment>
    )
  }
}

App.propTypes = {
  enqueueSnackbar: PropTypes.func.isRequired,
}

const MyApp = withSnackbar(App)

function IntegrationNotistack(
  { children, classes, color, variant, onClick },
  ...props
) {
  return (
    <SnackbarProvider maxSnack={1}>
      <MyApp />
    </SnackbarProvider>
  )
}

// export default withStyles(styles)(SoundSyncButton)
export default IntegrationNotistack
