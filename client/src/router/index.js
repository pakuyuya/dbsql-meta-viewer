import Vue from 'vue'
import Router from 'vue-router'
import TextSearch from '@/components/TextSearch'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'TextSearch',
      component: TextSearch
    }
  ]
})
