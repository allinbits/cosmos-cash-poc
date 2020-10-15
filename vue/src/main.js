import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import VueMaterial from "vue-material";
import "vue-material/dist/vue-material.min.css";
import "vue-material/dist/theme/default.css";

import IssuerList from "./components/IssuerList.vue";
import TokenAction from "./components/TokenAction.vue";

Vue.use(VueMaterial);
Vue.component("issuer-list", IssuerList);
Vue.component("token-action", TokenAction);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
