<template>
  <div class="home">
    Home
  </div>
</template>

<script>
export default {
  name: 'Home',
  created() {
    if (this.$store.getters['auth/loggedIn'])
      this.$router.push({name:"messages"})
    else
      this.$router.push({name:"account-login"})
  }
}
</script>
