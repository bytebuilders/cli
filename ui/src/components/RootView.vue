<template>
  <div class="bb-installer">
    <root-navbar />
    <ui-builder
      :show-options-step="false"
      :on-valid="doAction"
      :on-cancel="onCancel"
    />
  </div>
</template>
<script lang="ts" setup>
import RootNavbar from "./RootNavbar.vue";
import $axios from "../plugins/axios";
import { onMounted, provide, computed } from "vue";
import { uiBuilderJsonSetter } from "../composable/ui-builder";
import { useStore } from "./../store";

// Get ui-builder Value using computed property
const store = useStore();
const uiBuilderValue = computed(() => store.getters["wizard/model"]);

provide("$axios", $axios);

onMounted(() => {
  uiBuilderJsonSetter();
});

const doAction = () => {
  console.log("Called");
  const myTarget = JSON.parse(JSON.stringify(uiBuilderValue.value));
  console.log(myTarget);
};

const onCancel = () => {
  console.log("Click on clancel");
};
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
