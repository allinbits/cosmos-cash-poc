<template>
  <div>
    <md-card>
      <md-card-header class="md-layout">
        <div class="md-layout-item">
          <div class="md-title">Issuer List</div>
          <div class="md-subhead">Each issuer has a trading pool</div>
        </div>
        <md-button
          class="md-raised md-primary button_spacing"
          v-on:click="getIssuers"
          ><md-icon>refresh</md-icon></md-button
        >
      </md-card-header>

      <md-card-content>
        <br />
        <md-table>
          <md-table-row>
            <md-table-head>Name</md-table-head>
            <md-table-head>Address</md-table-head>
            <md-table-head>Token</md-table-head>
            <md-table-head>State</md-table-head>
            <md-table-head>Fee</md-table-head>
          </md-table-row>

          <md-table-row v-for="issuer in issuers" :key="issuer.name">
            <md-table-cell>{{ issuer.name }}</md-table-cell>
            <md-table-cell class="address">{{ issuer.address }}</md-table-cell>

            <md-table-cell>
              <md-menu>
                <md-button
                  md-menu-trigger
                  v-on:click="getTokens(issuer.address)"
                >
                  <md-icon>request_page</md-icon>
                </md-button>

                <md-menu-content>
                  <div v-for="token in tokens[issuer.address]">
                    <md-menu-item :key="token.denom">
                      <div>{{ token.denom }} - {{ token.amount }}</div>
                    </md-menu-item>
                  </div>
                </md-menu-content>
              </md-menu>
            </md-table-cell>
            <md-table-cell class="address">{{ issuer.state }}</md-table-cell>
            <md-table-cell class="address">{{ issuer.fee }}</md-table-cell>
          </md-table-row>
        </md-table>
      </md-card-content>
    </md-card>
  </div>
</template>

<style scoped>
.button_spacing {
  margin-right: 30px;
  margin-top: 5px;
}
</style>

<script>
export default {
  async created() {
    this.$store.dispatch("getIssuers", { val: "val" });
  },
  computed: {
    issuers() {
      return this.$store.state.issuers;
    },
    tokens() {
      return this.$store.state.tokens;
    },
  },
  methods: {
    async getIssuers() {
      this.$store.dispatch("getIssuers", { val: "val" });
    },
    async getTokens(address) {
      this.$store.dispatch("getTokens", { address });
    },
  },
};
</script>
