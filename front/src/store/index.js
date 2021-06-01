import Vue from 'vue'
import Vuex from 'vuex'
import snackbar from './modules/snack-bar'
import auth from './modules/auth'
import data from './modules/data'
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  getters:{
    locales : state => state.locales
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    snackbar,
    auth,
    data
  }
})
