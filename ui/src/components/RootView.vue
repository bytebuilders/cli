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
import { useRouter } from "vue-router";

//provide axios instance
provide("$axios", $axios);

//Init router & store
const router = useRouter();
const store = useStore();

//Check If model value available
const isModelValueAvailable: boolean = store.getters["getIsDataAvailable"];

//Set ui-builder initial value
onMounted(() => {
  if (!isModelValueAvailable) {
    uiBuilderJsonSetter();
  }
});

// Get ui-builder Value using computed property
const uiBuilderValue = computed(() => store.getters["wizard/model"]);

const doAction = async () => {
  try {
    const reqBody = JSON.parse(JSON.stringify(uiBuilderValue.value));
    const resp = await $axios.post("/apis/install", reqBody);
    const data = resp.data || "";
    store.commit("setMarkDown", data);
    store.commit("setIsDataAvailable", true);
    router.push("/info");
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
  margin-bottom: 30px;
  .ac-system-body {
    grid-template-columns: auto;
    .ui-builders-wrapper {
      margin: auto;
    }
  }
}
</style>
