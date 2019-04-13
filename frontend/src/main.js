import Vue from 'vue'
import App from './App.vue'
import router from './router'

Vue.config.productionTip = false


const API = 'http://localhost:5000';

const API_URL = API+'/api/';

Vue.http.options.root = API_URL;

new Vue({
  router,
  render: function (h) { return h(App) }
}).$mount('#app')
