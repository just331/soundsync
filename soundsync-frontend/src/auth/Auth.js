import auth0 from 'auth0-js'

export default class Auth {
  auth0 = new auth0.WebAuth({
    domain: process.env.AUTH0_DOMAIN,
    clientID: process.env.AUTH0_CLIENT_ID,
    redirectUri: 'http://localhost:3000/callback',
    responseType: 'token id_token',
    scope: 'openid',
    penis,
  })

  login() {
    this.auth0.authorize()
  }
}
