<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
import 'ress'

export default {
  name: 'App'
}
</script>

<style lang="scss">
@import "@/scss/common.scss";

html {
    height: 100%;
    overflow: hidden;
}

body {
    height: 100%;
}

#app {
  @include grow-box;
}
</style>
