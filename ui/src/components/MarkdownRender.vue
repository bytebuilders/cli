<template>
  <div class="bb-installer">
    <root-navbar />
    {{ markDown }}
  </div>
</template>

<script lang="ts" setup>
import RootNavbar from "./RootNavbar.vue";
import { computed } from "vue";
import { useStore } from "./../store";

// Get markdown from the store
const store = useStore();
const markDown = computed(() => store.getters["getMarkDown"]);
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
