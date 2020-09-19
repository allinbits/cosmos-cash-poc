<template>
  <div>
    <div class="containerMenu">
      <h1><router-link class="menu" to="/">Home</router-link></h1>
      <h1><router-link class="menu" to="/issuer">Issuers</router-link></h1>
      <h1><router-link class="menu" to="/poa">Consensus</router-link></h1>
    </div>
    <router-view />
  </div>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap");

body {
  margin: 0;
  font-family: "Inter", "Helvetica", sans-serif;
}
.containerMenu {
  max-width: 900px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  box-sizing: border-box;
  padding-left: 40px;
  padding-right: 40px;
  box-sizing: border-box;
  display: flex;
  justify-content: space-around;
}
.menu {
  background: rgba(0, 0, 0, 0.03);
  text-decoration: none;
  padding: 0 1.5rem;
  white-space: nowrap;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  color: rgba(0, 0, 0, 0.25);
  text-transform: uppercase;
  transition: all 0.1s;
}
</style>

<script>
export default {
  created() {
    this.$store.dispatch("init");
  }
};
</script>
