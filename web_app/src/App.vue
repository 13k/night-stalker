<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "app",
  created() {
    this.$log.debug("created");
    this.$store.dispatch("heroes/fetch");
    this.$store.dispatch("liveMatches/watch");
    this.$store.dispatch("liveMatches/fetch");
  }
};
</script>

<style lang="scss">
body {
  margin: 0;
  font-family: "Roboto", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  background-color: rgb(223, 223, 223);
}

.monospace {
  font-family: "Roboto Mono", monospace;
}

a {
  text-decoration: none;
  color: #2d6eaf;

  &:visited {
    color: #0a3e72;
  }
}
</style>
