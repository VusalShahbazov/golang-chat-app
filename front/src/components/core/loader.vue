<template>
  <div class="text-center mt-10 ">
    <v-progress-circular
        indeterminate
        :color="color"
    ></v-progress-circular>
  </div>
</template>

<script>
export default {
  name: "loader",
  props:{
    color:{
      default :'primary'
    }
  }
}
</script>

<style scoped>

</style>