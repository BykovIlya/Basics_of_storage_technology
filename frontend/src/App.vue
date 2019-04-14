<template>
  <div id="app">
    <b-navbar toggleable="md" type="dark" variant="dark">
      <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>
      <b-collapse is-nav id="nav_collapse">
        <b-navbar-nav>
          <b-nav-item :to="{ name: 'Boxoffice'}">{{ $t("message.boxoffice.title") }}</b-nav-item>
          <b-nav-item :to="{ name: 'Directors'}">{{ $t("message.directors.title") }}</b-nav-item>
          <b-nav-item :to="{ name: 'Movies'}">{{ $t("message.movies.title") }}</b-nav-item>
          <b-nav-item :to="{ name: 'Studios'}">{{ $t("message.studios.title") }}</b-nav-item>
          <b-nav-item :to="{ name: 'Analytics'}">{{ $t("message.analytics.title") }}</b-nav-item>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
    <router-view/>
  </div>
</template>

<script>
  export default {
    name: 'App'
  }
</script>


<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
  }
</style>