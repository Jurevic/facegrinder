import router from '../router'
import axios from 'axios'

const API_URL = 'http://0.0.0.0:8080/'
const LOGIN_URL = API_URL + 'api/v1/auth/login'
const SIGNUP_URL = API_URL + 'api/v1/users/'

export default {
  user: {
    authenticated: false
  },

  login_or_signup (action, context, creds, redirect) {
    let url = ''
    switch (action) {
      case 'login':
        url = LOGIN_URL
        break
      case 'signup':
        url = SIGNUP_URL
        break
      default:
        url = LOGIN_URL
    }
    axios.post(url, creds)
      .then((response) => {
        localStorage.setItem('api_token', response.data.api_token)

        this.user.authenticated = true

        if (redirect) {
          router.push(redirect)
        }
      }).catch((err) => {
        context.error = err.response.data
      })
  },

  logout () {
    localStorage.removeItem('api_token')
    this.user.authenticated = false
  },

  checkAuth () {
    const jwt = localStorage.getItem('api_token')
    if (jwt) {
      this.user.authenticated = true
    } else {
      this.user.authenticated = false
    }
  },

  getAuthHeader () {
    return {
      'Authorization': localStorage.getItem('api_token')
    }
  }
}
