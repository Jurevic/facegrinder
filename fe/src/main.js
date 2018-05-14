// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import axios from 'axios'
import VueAxios from 'vue-axios'
import Vue from 'vue'
import App from './App'
import router from './router'
import Vuetify from 'vuetify'
import VueVideoPlayer from 'vue-video-player'
import 'video.js/dist/video-js.css'
import 'videojs-flash'
import 'vuetify/dist/vuetify.min.css'
import { store } from './store'

Vue.use(Vuetify)
Vue.use(VueAxios, axios)
Vue.use(VueVideoPlayer)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})

// Request interceptor
axios.interceptors.request.use(function (config) {
  const token = store.state.jwt

  if (token != null) {
    config.headers.Authorization = `${token}`
  }

  return config
}, function (err) {
  return Promise.reject(err)
})
