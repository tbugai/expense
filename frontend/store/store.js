import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    expenses: []
  },

  getters: {
    getExpenses: state => {
      return state.expenses
    }
  },

  mutations: {
    'set-expenses': (state, expenses) => {
      state.expenses = expenses
    }
  },

  actions: {
    'fetch-expenses': ({commit}) => {
      return axios.get('/api/expenses')
        .then((response) => {
          commit('set-expenses', response.data)
        })
    }
  }
})

export default store
