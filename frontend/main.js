import Vue from 'vue'
import VueRouter from 'vue-router'
import router from './router'
import App from './App.vue'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.css'

Vue.use(VueRouter)
Vue.use(Vuetify)

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
