import React, { Component, useState } from "react";
import TextField from '@material-ui/core/TextField';
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

function CreateParty({ classes }) {
  const [values, setValues] = useState({
    nickName: '',
    phoneNumber: '',
    partyName: '',
  });

  const handleChange = name => event => {
    setValues({ ...values, [name]: event.target.value });
  }

  return (
    <div>
      <form>
        <TextField
          id="partyName"
          label="Party Name"
          className={classes.textField}
          value={values.partyName}
          onChange={handleChange('partyName')}
        />
        <TextField
          id="phoneNumber"
          label="Phone Number"
          className={classes.textField}
          value={values.phoneNumber}
          onChange={handleChange('phoneNumber')}
        />
        <TextField
          id="nickName"
          label="Nick Name"
          className={classes.textField}
          value={values.nickName}
          onChange={handleChange('nickName')}
        />
      </form>
    </div>
  )
}

export default withStyles(styles)(CreateParty);
