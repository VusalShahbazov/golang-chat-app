<template>
  <div class="fill-height bgMainImage primary" ref="bgMainImage">
    <guest-header/>
    <AnimatedRouterView/>
  </div>
</template>

<script>
import GuestHeader from "../../components/core/guest-header";
import AnimatedRouterView from "../../components/core/animated-router-view";

export default {
  name: "Layout",
  components: {GuestHeader, AnimatedRouterView},
  mounted() {
    this.selectImageForBackground()
  },
  methods: {
    selectImageForBackground() {
      const num = Math.floor(Math.random() * 9)
      this.$refs.bgMainImage.style.backgroundImage = `url("/images/background/${num}.png")`
    }
  }
}
</script>
