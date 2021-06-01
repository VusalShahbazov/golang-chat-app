<template>
  <div class="bg-light-blue bgMainImage fill-height" ref="bgMainImage">
    <user-header />
    <animated-router-view :key="freelancer" />
  </div>
</template>

<script>
import AnimatedRouterView from "../../components/core/animated-router-view";
import UserHeader from "../../components/core/user-header";
export default {
  name: "Layout",
  components: {UserHeader, AnimatedRouterView},
  computed:{
    freelancer(){
      return this.$store.getters['auth/is_freelancer'];
    }
  },
  mounted() {
    this.selectImageForBackground()
  },
  methods:{
    selectImageForBackground(){
      const num = Math.floor(Math.random() * 9)
      this.$refs.bgMainImage.style.backgroundImage = `url("/images/background/${num}.png")`
    }
  }
}
</script>

<style scoped>

</style>