import React, { useState } from "react";
import {
    Route,
    NavLink,
} from "react-router-dom";
import Party from "./Party";
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';

const styles = (theme) => {
  return ({
    container: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    textField: {
      marginLeft: theme.spacing.unit,
      marginRight: theme.spacing.unit,
      width: 200,
    },
    dense: {
      marginTop: 19,
    },
    menu: {
      width: 200,
    },
  })
};

function JoinParty({ classes }) {
  const [values, setValues] = useState({
    partyCode: '',
    nickName: '',
  });

  const handleChange = name => event => {
    setValues({ ...values, [name]: event.target.value });
  }

  // TODO: Make a post to join the party then redirect the user to the join party screen
  // TODO: Check to see if they have a valid jwt token first (means they are logged in)
  const handleJoinParty = e => {
    e.preventDefault();
    console.log("Joined Party!");

  }

  return (
    <div>
      <Button variant="contained" color="primary"><NavLink to="/CreateParty">Create Party</NavLink></Button>
      <Button variant="contained" color="primary"><NavLink to="/">Join Party</NavLink></Button>
      <Grid item sm={6}>
        <TextField
          id="partyCode"
          label="Party Code"
          required
          className={classes.textField}
          value={values.partyCode}
          onChange={handleChange('partyCode')}
        />
      </Grid>
      <Grid item sm={6}>
        <TextField
          id="nickName"
          label="Nickname"
          required
          className={classes.textField}
          value={values.nickName}
          onChange={handleChange('nickName')}
        />
      </Grid>
      <Grid item sm={6}>
        <Route path="/Party" component={Party}/>
        <Button variant="contained" color="primary" type="submit" onClick={handleJoinParty}><NavLink to="/Party">Join Party</NavLink></Button>
      </Grid>
    </div>
  )
}

export default withStyles(styles)(JoinParty);
