import React, { useState } from "react";
import {
  Button,
  BottomNavigation,
  BottomNavigationAction,
} from '@material-ui/core/';
import RepeatIcon from '@material-ui/icons/Repeat';
import SkipPreviousIcon from '@material-ui/icons/SkipPrevious';
import PauseIcon from '@material-ui/icons/Pause';
import SkipNextIcon from '@material-ui/icons/SkipNext';
import ShuffleIcon from '@material-ui/icons/Shuffle';
import { withStyles } from '@material-ui/core/styles';

const styles = (theme) => {
  return ({
    musicControls: {
      position: 'fixed',
      bottom: 0,
    },
  })
};

function MusicControl({ classes }) {

  // TODO: Change these menu icons to actual music player icons
  return (
    <div>
      <BottomNavigation className={classes.musicControls}>
          <BottomNavigationAction label="Repeat" icon={<RepeatIcon />} />
          <BottomNavigationAction label="Previous" icon={<SkipPreviousIcon />} />
          <BottomNavigationAction label="Pause" icon={<PauseIcon />} />
          <BottomNavigationAction label="Next" icon={<SkipNextIcon />} />
          <BottomNavigationAction label="Shuffle" icon={<ShuffleIcon />} />
      </BottomNavigation>
    </div>
  )
}

export default withStyles(styles)(MusicControl);
