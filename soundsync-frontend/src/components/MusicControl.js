import React from 'react'
import { BottomNavigation, BottomNavigationAction } from '@material-ui/core/'
import RepeatIcon from '@material-ui/icons/Repeat'
import SkipPreviousIcon from '@material-ui/icons/SkipPrevious'
import PauseIcon from '@material-ui/icons/Pause'
import SkipNextIcon from '@material-ui/icons/SkipNext'
import ShuffleIcon from '@material-ui/icons/Shuffle'
import { withStyles } from '@material-ui/core/styles'

const styles = (theme) => {
  return {
    musicControls: {
      position: 'fixed',
      bottom: 0,
    },
  }
}

function MusicControl({ classes }) {
  const handleRepeat = (e) => {
    e.preventDefault()
    console.log('Repeat Clicked!')
  }

  const handlePrevious = (e) => {
    e.preventDefault()
    console.log('Previous Clicked!')
  }

  const handlePause = (e) => {
    e.preventDefault()
    console.log('Pause Clicked!')
  }

  const handleNext = (e) => {
    e.preventDefault()
    console.log('Next Clicked!')
  }

  const handleShuffle = (e) => {
    e.preventDefault()
    console.log('Shuffle Clicked!')
  }

  // TODO: Change these menu icons to actual music player icons
  return (
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        alignContent: 'center',
      }}
    >
      <BottomNavigation className={classes.musicControls}>
        <BottomNavigationAction
          onClick={handleRepeat}
          label='Repeat'
          icon={<RepeatIcon />}
        />
        <BottomNavigationAction
          onClick={handlePrevious}
          label='Previous'
          icon={<SkipPreviousIcon />}
        />
        <BottomNavigationAction
          onClick={handlePause}
          label='Pause'
          icon={<PauseIcon />}
        />
        <BottomNavigationAction
          onClick={handleNext}
          label='Next'
          icon={<SkipNextIcon />}
        />
        <BottomNavigationAction
          onClick={handleShuffle}
          label='Shuffle'
          icon={<ShuffleIcon />}
        />
      </BottomNavigation>
    </div>
  )
}

export default withStyles(styles)(MusicControl)
