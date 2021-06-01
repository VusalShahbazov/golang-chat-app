<template>
  <v-app class="fill-height">
    <animated-router-view/>
    <v-snackbar
        min-width="250px"
        width="50px"
        bottom
        right
        :color="snack.color"
        v-model="snack.show">
      <span>
      {{ snack.text }}
      </span>
      <template #action>
        <v-btn text icon @click="$store.dispatch('snackbar/HIDE_SNACK_BAR')">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </template>
    </v-snackbar>
  </v-app>
</template>
<script>
import AnimatedRouterView from "./components/core/animated-router-view";

export default {
  name:"App",
  components: {AnimatedRouterView},

  mounted() {
    document.getElementById("main-loader").style.display = 'none'

  },
  computed: {
    snack() {
      return this.$store.getters['snackbar/snack']
    }
  },

}
</script>
<style>
* {
  font-family: "Open Sans", serif !important;
}

.btn {
  padding: 8px 16px !important;
  letter-spacing: normal !important;
  text-transform: none !important;
}

.letter-spacing-normal {
  letter-spacing: normal !important;
}

.bgMainImage {
  /*background-image: url("../../../public/images/download.jpg");*/
  background-position: center center;
  background-attachment: fixed;
  background-size: cover;
}

.bg-light-blue {
  background-color: #EDF0F6;
}

.v-toolbar__content {
  padding-top: 0 !important;
  padding-bottom: 0 !important;
}

.btn-pink {
  background: linear-gradient(94.65deg, #FF2D64 0%, #D00037 100%), #FFFFFF;
}

.pink-color {
  color: #E6154C !important;
}

.btn-shaped {
  border-radius: 14px 0 14px 0 !important;
}

.blue-color {
  color: #3f2bfa !important;
}

.border-radius-default {
  border-radius: 4px !important;
}

.font14 {
  font-size: 14px !important;
}

.font13 {
  font-size: 13px !important;
}

.font11 {
  font-size: 11px !important;
}

.light-blue-tab {
  background-color: #eeedff !important;
}

.card-title {
  font-size: 18px !important;
  font-weight: bold;
  line-height: 23px !important;
}

.card-description {
  font-size: 13px !important;
  line-height: 20px !important;
}
/* width */
::-webkit-scrollbar {
  width: 5px;
}

/* Track */
::-webkit-scrollbar-track {
  background: #f1f1f1;
}

/* Handle */
::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 2px;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: #919191;
}

</style>
