import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import './registerServiceWorker'
import store from './store'


store.dispatch("auth/GET_PROFILE")
    .catch(err => {
        console.warn(err)
    })
    .finally(() => {

        Vue.config.productionTip = false

        new Vue({
            vuetify,
            router,
            store,
            render: h => h(App)
        }).$mount('#app')
    })

