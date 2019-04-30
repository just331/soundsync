import React from 'react'
import {
  Button,
  IconButton,
  Typography,
  Toolbar,
  AppBar,
} from '@material-ui/core'
import MenuIcon from '@material-ui/icons/Menu'
import { withStyles } from '@material-ui/core/styles'
import { MuiThemeProvider, createMuiTheme, withTheme } from '@material-ui/core/styles'

const styles = createMuiTheme({
  palette: {
    primary: {
      main: '#fff',
    },
    secondary: {
      main: '#3CBEAF',
    },
  },
})

export default class AddSong extends React.Component{
  constructor(props){
    super(props);

    this.state = {
      song: ''
    };

    this.addSong = this.addSong.bind(this);
  }
  addSong = (e) => {
    e.preventDefault();
    //get the fruit object name from the form
    var song = this.refs.songName.value;
    //if we have a value
    //call the addFruit method of the App component
    //to change the state of the fruit list by adding an new item
    if(typeof song === 'string' && song.length > 0) {
      this.props.queueSong(song);
      //reset the form
      this.refs.songForm.reset();
    }
   }
   render() {
    return(
      <form className={styles} ref="songForm" onSubmit={this.addSong}>
      <div className="form-group">
        <input type="text" id="songItem" placeholder="Lemon" ref="songName" className="form-control" />
      </div>
      <Button type="submit" className={styles.button}>Add Song </Button>
     </form>
    )
   }
};