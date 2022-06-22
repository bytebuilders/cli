import { createApp } from "vue";
import App from "./App.vue";

// @ts-ignore
import "./assets/styles/scss/main.scss";

// vue router
import router from "./router";

//vuex
import { store, key } from "./store";

//ui-builder
import i18n from "./plugins/i18n";
import UiBuilder from "@appscode/ui-builder";

//monaco-editor
import { useMonaco } from "./plugins/monaco";
useMonaco();

createApp(App)
  .use(router)
  .use(store, key)
  .use(i18n)
  .use(UiBuilder, {
    store,
    key,
    storeOptions: { preserveState: false },
    i18n: i18n.global,
  })
  .mount("#app");
