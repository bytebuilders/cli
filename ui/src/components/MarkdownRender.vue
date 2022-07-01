<template>
  <div class="bb-installer">
    <root-navbar />
    <div v-html="mdToHtml"></div>
  </div>
</template>

<script lang="ts" setup>
import RootNavbar from "./RootNavbar.vue";
import { computed } from "vue";
import { useStore } from "./../store";
import { marked } from "marked";

// Get markdown from the store
const store = useStore();
const markDown = computed(() => store.getters["getMarkDown"]);

const mdToHtml = marked.parse(markDown.value);
</script>
<style lang="scss">
.bb-installer {
  .ac-system-body {
    grid-template-columns: auto;
    .ui-builders-wrapper {
      margin: auto;
    }
  }
}
</style>
