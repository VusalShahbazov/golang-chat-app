<template>
  <v-app-bar
      color="white"
      height="60px"
  >
    <v-toolbar-title class="d-flex align-center">
      <img
          :src="$vuetify.breakpoint.smAndDown  ?'/images/golang.png' : '/images/golang.png'"
          height="40px">
    </v-toolbar-title>

    <v-spacer></v-spacer>
  </v-app-bar>
</template>

<script>
export default {
  name: "guest-header",

}
</script>

<style scoped>

</style>