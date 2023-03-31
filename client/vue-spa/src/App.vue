<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <Yamabiko/>
</template>

<script>
import Yamabiko from './components/Yamabiko.vue'

export default {
  name: 'App',
  components: {
    Yamabiko
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
