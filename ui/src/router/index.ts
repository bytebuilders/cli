import { createRouter, createWebHistory } from "vue-router";
import RootView from "../components/RootView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: RootView,
    },
  ],
});

export default router;
