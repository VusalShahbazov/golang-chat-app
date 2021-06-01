<template>
  <v-card flat color="transparent" class="mt-10">
    <v-card-text class="d-flex justify-center">
      <slot name="icon">
        <v-icon size="200" >{{ icon }}</v-icon>
      </slot>
    </v-card-text>
    <v-card-text class="text-center text-h4 black--text font-weight-bold">
      {{ this.text }}
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "not-found",
  props: {
    icon: {
      type :String ,
      default : "mdi-magnify-close"
    },
    text :{
      type :String,
      default : "No Projects"
    }
  },

  
}
</script>

<style scoped>

</style>