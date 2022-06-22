import vue from "@vitejs/plugin-vue";
import { defineConfig } from "vite";
import pluginYaml from "vite-plugin-yaml2";

import { join } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), pluginYaml()],
  resolve: {
    alias: {
      "@": join(__dirname, "/src"),
      "vue-i18n": "vue-i18n/dist/vue-i18n.cjs.js",
    },
  },
  server: {
    port: 5992,
  },
});
