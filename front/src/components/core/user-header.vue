<template>
  <div>

  <v-app-bar
      color="white custom-toolbar"
      elevation="10"
      :elevate-on-scroll="true"
      class="pt-0 pb-0"
      height="55px"
  >
    <v-toolbar-title class="d-flex align-center ">

      <img
          :src="$vuetify.breakpoint.smAndDown  ?'/images/golang.png' : '/images/golang.png'"
          height="40px">

    </v-toolbar-title>

    <v-spacer></v-spacer>
    <v-app-bar-nav-icon
        v-if="$vuetify.breakpoint.smAndDown"
        @click="drawer = !drawer">
      <v-icon>mdi-menu</v-icon>
    </v-app-bar-nav-icon>
  </v-app-bar>
    <v-navigation-drawer
        v-if="$vuetify.breakpoint.smAndDown"
        app
        temporary
        v-model="drawer"
    >
      <rooms @selectRoom="selectRoom" />

    </v-navigation-drawer>
    </div>
</template>

<script>

import Rooms from "./rooms";
export default {
  name: "guest-header",
  components: {Rooms},
  computed: {
    user(){
      return this.$store.getters['auth/user']
    }
  },
  data() {
    return {
      drawer:false,
    }
  },
  methods:{
    selectRoom(room){
      this.drawer = false
      this.$emit('selectRoom' , room)
    }
  }

}
</script>

<style>

</style>