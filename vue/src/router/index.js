import Vue from "vue";
import VueRouter from "vue-router";
import Index from "../views/Index.vue";
import Issuer from "../views/Issuer.vue";
import Admin from "../views/Admin.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Index,
  },
  {
    path: "/issuers",
    component: Issuer,
  },
  {
    path: "/admin",
    component: Admin,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
