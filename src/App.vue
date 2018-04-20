<template>
  <div id="app">
    <!-- <img src="./assets/logo.png"> -->
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
html,body {
  margin: 0;
  padding: 0;
  color: #333;
}
#app {
  font-family: PingFang SC,'Avenir', Helvetica Neue,Helvetica, Arial, sans-serif;
}
</style>
