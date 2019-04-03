import React, { Component, useState } from "react";
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

function CreateParty({ classes }) {
  const [values, setValues] = useState({
    nickName: '',
    phoneNumber: '',
    partyName: '',
    verifyCode: '',
    isVerifyCodeSent: false,
  });

  const handleChange = name => event => {
    setValues({ ...values, [name]: event.target.value });
  }

  // TODO: When button is clicked POST data and verify if all fields are valid
  const handleSendVerify = e => {
    e.preventDefault();
    setValues({ ...values, isVerifyCodeSent: true });
    console.log("Verification code sent!");
  }

  // TODO: Make a POST to see if the verification code is correct before creating the party
  const handleCreateParty = e => {
    e.preventDefault();
    console.log("Created a party!");
  }

  return (
    <div>
      <form>
        <Grid item sm={6}>
          <TextField
            id="partyName"
            label="Party Name"
            required
            className={classes.textField}
            value={values.partyName}
            onChange={handleChange('partyName')}
          />
        </Grid>
        <Grid item sm={6}>
          <TextField
            id="phoneNumber"
            label="Phone Number"
            required
            className={classes.textField}
            value={values.phoneNumber}
            onChange={handleChange('phoneNumber')}
          />
        </Grid>
        {values.isVerifyCodeSent ? (
          <div>
            <Grid item sm={6}>
              <TextField
                id="verifyCode"
                label="Enter Verification Code"
                required
                className={classes.textField}
                value={values.verifyCode}
                onChange={handleChange('verifyCode')}
              />
            </Grid>
            <Grid item sm={6}>
              <Button variant="contained" color="primary" type="submit" onClick={handleCreateParty}>Create Party</Button>
            </Grid>
          </div>
        ) : (
          <Grid item sm={6}>
            <Button variant="contained" color="primary" type="submit" onClick={handleSendVerify}>Send Verification Code</Button>
          </Grid>
        )}
      </form>
    </div>
  );
}

export default withStyles(styles)(CreateParty);
