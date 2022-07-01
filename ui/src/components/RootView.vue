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

const doAction = async () => {
  try {
    const reqBody = JSON.parse(JSON.stringify(uiBuilderValue.value));
    const resp = await $axios.post("/apis/install", reqBody);
    const data = resp.data || "";
    store.commit("setMarkDown", data);
    console.log(data);
  } catch (error) {
    console.log(error);
  }
};

const onCancel = () => {
  window.location.href = "/";
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
