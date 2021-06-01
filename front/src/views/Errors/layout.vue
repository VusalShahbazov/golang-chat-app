<template>
  <div class="fill-height bgMainImage primary" ref="bgMainImage">
    <guest-header/>
    <AnimatedRouterView/>
  </div>
</template>

<script>
import AnimatedRouterView from "../../components/core/animated-router-view";
import GuestHeader from "../../components/core/guest-header";
export default {
  name: "Layout",
  components: {GuestHeader, AnimatedRouterView},
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

<style>
.bgErrorImage{
  /*background-image: url("../../../public/images/download.jpg");*/
  background-position: center center;
  background-attachment: fixed;
  background-size: cover;
  /*background-color: #464646;*/
}
</style>