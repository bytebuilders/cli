<template>
  <div class="bb-installer">
    <root-navbar />
    <div v-html="mdToHtml"></div>
    <button @click="onBackClick">Back</button>
  </div>
</template>

<script lang="ts" setup>
import RootNavbar from "./RootNavbar.vue";
import { computed } from "vue";
import { useStore } from "./../store";
import { useRouter } from "vue-router";
import { marked } from "marked";

//init store & router
const store = useStore();
const router = useRouter();

// Get markdown from the store
const markDown = computed(() => store.getters["getMarkDown"]);

const onBackClick = () => {
  router.push("./");
};

const mdToHtml = marked.parse(markDown.value);
</script>
<style lang="scss" scoped>
.bb-installer {
  margin-top: 50px;
}
</style>
