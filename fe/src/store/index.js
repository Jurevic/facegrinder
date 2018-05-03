import Vue from 'vue'
import Vuex from 'vuex'
import jwtDecode from 'jwt-decode'
import axios from 'axios'
import router from '../router'

Vue.use(Vuex)

export const store = new Vuex.Store({
  state: {
    appTitle: 'FaceGrinder',
    user: null,
    jwt: localStorage.getItem('t'),
    error: null,
    loading: false,
    endpoints: {
      obtainJWT: 'http://0.0.0.0:8080/api/auth/login',
      refreshJWT: 'http://0.0.0.0:8080/api/auth/refresh',
      users: 'http://0.0.0.0:8080/api/v1/users/',
      faces: 'http://0.0.0.0:8080/api/v1/faces/',
      channels: 'http://0.0.0.0:8080/api/v1/channels/',
      processors: 'http://0.0.0.0:8080/api/v1/processors/'
    }
  },
  mutations: {
    setUser (state, payload) {
      state.user = payload
    },
    setToken (state, token) {
      localStorage.setItem('t', token)
      state.jwt = token
    },
    clearToken (state) {
      localStorage.removeItem('t')
      state.jwt = null
    },
    setError (state, payload) {
      state.error = payload
    },
    setLoading (state, payload) {
      state.loading = payload
    }
  },
  actions: {
    refreshToken ({commit, state}) {
      const payload = {
        token: state.jwt
      }

      axios.post(this.state.endpoints.refreshJWT, payload)
        .then((response) => {
          commit('updateToken', response.data.token)
        })
        .catch((error) => {
          console.log(error)
        })
    },
    inspectToken ({actions}) {
      const token = this.state.jwt
      if (token) {
        const decoded = jwtDecode(token)
        const exp = decoded.exp
        const iat = decoded.iat

        if (exp - (Date.now() / 1000) < 1800 && (Date.now() / 1000) - iat < 628200) {
          this.dispatch('refreshToken')
        } else if (exp - (Date.now() / 1000) < 1800) {
          // DO NOTHING, DO NOT REFRESH
        } else {
          actions.userSignOut()
          router.push('/signin')
        }
      }
    },
    userSignUp ({commit, state}, payload) {
      commit('setLoading', true)
      axios.post(state.endpoints.users, payload)
        .then((response) => {
          commit('setLoading', false)
          commit('setError', null)
          router.push('/signin')
        })
        .catch((error) => {
          commit('setError', error.message)
          commit('setLoading', false)
        })
    },
    userSignIn ({commit, state}, payload) {
      commit('setLoading', true)
      axios.post(state.endpoints.obtainJWT, payload)
        .then((response) => {
          commit('setToken', response.data.token)
          commit('setUser', {email: 'test'})
          commit('setLoading', false)
          commit('setError', null)
          router.push('/home')
        })
        .catch((error) => {
          commit('setError', error.message)
          commit('setLoading', false)
        })
    },
    userSignOut ({commit}) {
      commit('setUser', null)
      commit('clearToken')
      router.push('/')
    }
  },
  getters: {
    isAuthenticated (state) {
      return state.user !== null && state.user !== undefined && state.jwt !== null && state.jwt !== undefined
    }
  }
})
