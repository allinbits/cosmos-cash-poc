<template>
  <div>
    <md-card class="z md-elevation-5 sp-container">
      <div class="md-layout">
        <md-tabs md-sync-route>
          <div class="md-layout-item">
            <md-tab
              id="tab-home"
              md-label="Home"
              md-icon="home"
              to="/"
              exact
            ></md-tab>
            <md-tab
              id="tab-Id"
              md-label="Identity"
              md-icon="perm_identity"
              to="/ids"
              exact
            ></md-tab>
            <md-tab
              id="tab-issuer"
              md-label="Issuers"
              md-icon="supervisor_account"
              to="/issuers"
            ></md-tab>
            <md-tab
              id="tab-admin"
              md-icon="admin_panel_settings"
              md-label="admin"
              to="/admin"
            ></md-tab>
            <md-tab
              id="tab-poa"
              md-icon="vpn_lock"
              md-label="validators"
              to="/validators"
            ></md-tab>
          </div>
        </md-tabs>
        <sp-sign-in class="login-bar md-layout-item" />
      </div>
    </md-card>

    <router-view />
  </div>
</template>

<style>
.sp-container {
  margin: 0 auto;
  margin-top: 1rem;
  margin-bottom: 50px;
  max-width: 1200px;
}
.z {
  z-index: 100 !important;
}
.login-bar {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  margin-right: 10px;
}
</style>

<script>
import * as sp from "@tendermint/vue";
export default {
  components: { ...sp },
  created() {
    this.$store.dispatch("cosmos/init");
  },
};
</script>
