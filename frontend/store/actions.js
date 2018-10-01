import axios from 'axios'

export default {
  'fetch-expenses': ({commit}) => {
    return axios.get('/api/expenses')
      .then((response) => {
        commit('set-expenses', response.data)
      })
  }
}
