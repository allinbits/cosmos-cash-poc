<template>
  <div>
    <div class="containerMenu">
    <md-tabs md-sync-route>
      <md-tab id="tab-home" class="menu_button" md-label="Home" md-icon="home" to="/" exact></md-tab>
      <md-tab id="tab-issuer" class="menu_button" md-label="Issuers" md-icon="supervisor_account" to="/issuer"></md-tab>
      <md-tab id="tab-poa" class="menu_button" md-icon="vpn_lock" md-label="Validators" to="/poa"></md-tab>
    </md-tabs>
    </div>
    <router-view />
  </div>
</template>

<style>

body {
  margin: 0;
  font-family: "Inter", "Helvetica", sans-serif;
}
.containerMenu {
  max-width: 1200px;
  margin-left: auto;
  margin-right: auto;
  padding-left: 40px;
  padding-right: 40px;
  display: flex;
}
.menu_button {
width: 100px;
}
</style>

<script>
export default {
  created() {
    this.$store.dispatch("init");
  }
};
</script>
