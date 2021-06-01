<template>
  <v-tooltip top :open-on-hover="true" v-model="show" :open-on-click="true">
    <template v-slot:activator="{ on, attrs }">
      <v-icon
          tabindex="-1"
          size="20"
          :style="{'position':'relative' , 'top' : top}"
          @click="show = !show"
          class="color-black"
          v-bind="attrs"
          v-on="on"
      >mdi-help-circle</v-icon
      >
    </template>
    <span v-html="text"></span>
  </v-tooltip>
</template>

<script>
export default {
  name: "vHelp",
  props: ["text", "top"],
  data() {
    return {
      show: false
    };
  }
};
</script>
