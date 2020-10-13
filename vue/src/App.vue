<template>
  <div>
    <md-tabs class="sp-container md-elevation-5" md-sync-route>
      <md-tab
        id="tab-home"
        md-label="Home"
        md-icon="home"
        to="/"
        exact
      ></md-tab>
      <md-tab
        id="tab-issuer"
        md-label="Issuers"
        md-icon="supervisor_account"
        to="/issuers"
      ></md-tab>
      <md-tab
        id="tab-poa"
        md-icon="vpn_lock"
        md-label="admin"
        to="/admin"
      ></md-tab>
    </md-tabs>
    <router-view />
  </div>
</template>

<style>
.sp-container {
  margin-top: 1rem;
  margin-bottom: 50px;
  max-width: 1200px;
}
</style>

<script>
export default {
  created() {
    this.$store.dispatch("cosmos/init");
  },
};
</script>
