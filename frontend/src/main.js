import Vue from 'vue'
import App from './App.vue'
import router from './router'
import i18n from '../localization/index.js'

import VueI18n from 'vue-i18n';
Vue.use(VueI18n);

import BootstrapVue from 'bootstrap-vue'
Vue.use(BootstrapVue);
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import VueResorce from 'vue-resource'
Vue.use(VueResorce);

import Vuex from 'vuex'
Vue.use(Vuex);

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { library } from '@fortawesome/fontawesome-svg-core'

import { faTrash } from '@fortawesome/free-solid-svg-icons'
import { faEdit } from '@fortawesome/free-solid-svg-icons'
import { faEye } from '@fortawesome/free-solid-svg-icons'
import { faPassport } from '@fortawesome/free-solid-svg-icons'
import { faTimesCircle } from '@fortawesome/free-solid-svg-icons'

library.add(faTrash,faEdit,faEye,faPassport,faTimesCircle);


Vue.component('font-awesome-icon', FontAwesomeIcon);

Vue.config.productionTip = false;


const API = 'http://localhost:5000';

const API_URL = API+'/api/';

Vue.http.options.root = API_URL;

new Vue({
  router,
  i18n,
  render: function (h) { return h(App) }
}).$mount('#app');
