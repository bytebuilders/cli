<template>
  <div>
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
