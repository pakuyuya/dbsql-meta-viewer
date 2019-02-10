// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import config from 'config'
import services from './services/entry'

Vue.config.productionTip = false

const configMixin = Vue.mixin({
  created: function () {
    this.$config = config

    for (const key in services) {
      var tagservice = '$' + key
      if (this[tagservice]) {
        console.error(`module name conflicted with $vue.${tagservice}! we take overwrite to ./services/modules.`)
      }
      this['$' + key] = services[key]
    }
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  configMixin,
  components: { App },
  template: '<App/>'
})
