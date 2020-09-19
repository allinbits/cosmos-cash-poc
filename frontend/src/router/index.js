import Vue from "vue";
import VueRouter from "vue-router";
import Index from "../views/Index.vue";
import Poa from "../views/Poa.vue";
import Issuer from "../views/Issuer.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Index,
  },
  {
    path: "/poa",
    component: Poa,
  },
  {
    path: "/issuer",
    component: Issuer,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
