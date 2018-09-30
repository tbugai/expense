import VueRouter from 'vue-router'
import DashboardComponent from './components/dashboard'

export default new VueRouter({
  routes: [
    { path: '/', component: DashboardComponent }
  ],
  mode: 'history'
})
