import Vue from 'vue'
import Router from 'vue-router'
import Directors from './components/Directors'
import Boxoffice from "./components/Boxoffice";
import Studios from "./components/Studios";
import Movies from "./components/Movies"
Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Boxoffice',
      component: Boxoffice
    },
    {
      path: '/directors',
      name: 'Directors',
      component: Directors
    },
    {
      path: '/movies',
      name: 'Movies',
      component: Movies
    },
    {
      path: '/studios',
      name: 'Studios',
      component: Studios
    },
  ]
})
