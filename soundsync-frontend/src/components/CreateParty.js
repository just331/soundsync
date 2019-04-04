import React, { useState } from 'react'
import { withStyles } from '@material-ui/core/styles'
import SoundSyncInput from 'components/Input'
import SoundSyncButton from 'components/Button'
import Grid from '@material-ui/core/Grid'
import AppContainer from 'components/AppContainer'
import Logo from 'components/Logo'

const styles = (theme) => {
  return {
    Container: {
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      padding: '135px 0px',
    },
    dense: {
      marginTop: 19,
    },
    menu: {
      width: 200,
    },
  }
}

function CreateParty({ classes }) {
  const [values, setValues] = useState({
    nickName: '',
    phoneNumber: '',
    partyName: '',
    verifyCode: '',
    isVerifyCodeSent: false,
  })

  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value })
  }

  // TODO: When button is clicked POST data and verify if all fields are valid
  const handleSendVerify = (e) => {
    e.preventDefault()
    setValues({ ...values, isVerifyCodeSent: true })
    console.log('Verification code sent!')
  }

  // TODO: Make a POST to see if the verification code is correct before creating the party
  const handleCreateParty = (e) => {
    e.preventDefault()
    console.log('Created a party!')
  }

  return (
    <AppContainer>
      <Grid container className={classes.Container}>
        <Logo />
        <Grid item sm={6}>
          <SoundSyncInput
            id='partyName'
            value={values.partyName}
            placeholder='Enter Party Name'
            onChange={handleChange('partyName')}
          />
        </Grid>
        <Grid item sm={6}>
          <SoundSyncInput
            id='phoneNumber'
            value={values.phoneNumber}
            placeholder='Enter Phone Number'
            onChange={handleChange('phoneNumber')}
          />
        </Grid>
        {values.isVerifyCodeSent ? (
          <>
            <Grid item sm={6}>
              <SoundSyncInput
                id='verifyCode'
                placeholder='Enter Verification Code'
                value={values.verifyCode}
                onChange={handleChange('verifyCode')}
              />
            </Grid>
            <Grid item sm={6}>
              <SoundSyncButton
                variant='contained'
                color='secondary'
                onClick={handleCreateParty}
              >
                Create Party
              </SoundSyncButton>
            </Grid>
          </>
        ) : (
          <Grid item sm={6}>
            <SoundSyncButton
              variant='contained'
              color='secondary'
              type='submit'
              onClick={handleSendVerify}
            >
              Send Verification Code
            </SoundSyncButton>
          </Grid>
        )}
      </Grid>
    </AppContainer>
  )
}

export default withStyles(styles)(CreateParty)
