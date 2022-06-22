import vue from "@vitejs/plugin-vue";
import { defineConfig } from "vite";

import { join } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
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
