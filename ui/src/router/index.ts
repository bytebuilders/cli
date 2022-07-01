import MarkDownRender from "../components/MarkdownRender.vue";
import RootView from "../components/RootView.vue";
import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: RootView,
    },
    {
      path: "/info",
      name: "info",
      component: MarkDownRender,
    },
  ],
});

export default router;
