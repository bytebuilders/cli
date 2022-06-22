<template>
  <div class="bb-installer">
    <ui-builder
      :show-options-step="false"
      :on-valid="doAction"
      :on-cancel="onCancel"
    />
  </div>
</template>
<script lang="ts" setup>
import $axios from "../plugins/axios";
import { onMounted, provide } from "vue";
import { uiBuilderJsonSetter } from "../composable/ui-builder";

provide("$axios", $axios);

onMounted(() => {
  uiBuilderJsonSetter();
});

const doAction = () => {
  console.log("Called Do action");
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
