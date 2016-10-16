// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Router from 'vue-router'

import App from './App'
import ResourceView from './components/ResourceView'
import SearchView from './components/SearchView'

Vue.use(Router)

const router = new Router({
  routes: [
    {name: 'search', path: '/search', component: SearchView},
    {name: 'resource', path: '/resources/:id', component: ResourceView}
  ]
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router: router,
  render: h => h(App)
})

