<template>
  <loader/>
</template>

<script>
import Loader from "../../../components/core/loader"
import validations from "../../../mixins/validate";
export default {
  name: "Logout",
  mixins:[validations],
  components: {Loader},
  beforeCreate() {
    this.$store.dispatch("auth/LOGOUT")
    .catch(err => {
      this.handleError(err)
    })
    .finally(()=>{
      this.$router.push("/")
    })
  }
}
</script>

<style scoped>

</style>