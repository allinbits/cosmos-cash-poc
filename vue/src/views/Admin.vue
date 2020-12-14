<template>
  <div class="sp-container">
    <div class="top_spacing">
      <md-card class="md-elevation-5">
        <md-card-header class="md-layout">
          <div class="md-layout-item">
            <div class="md-title">Admin actions</div>
            <div class="md-subhead">
              Each issuer can mint, burn and freeze tokens
            </div>
          </div>
        </md-card-header>
        <token-action type="Burn" />
        <token-action type="Mint" />
        <div v-model="issuer">
          <div v-if="issuer">
            <div v-if="issuer.state === 'accepted'">
              <token-action type="Freeze" />
            </div>
            <div v-if="issuer.state === 'frozen'">
              <token-action type="Unfreeze" />
            </div>
          </div>
          <div v-else>
            <user-action type="Withdraw" />
          </div>
        </div>
      </md-card>
    </div>
  </div>
</template>

<style scoped>
.top_spacing {
  margin-top: 50px;
}
</style>

<script>
import * as sp from "@tendermint/vue";
export default {
  components: { ...sp },
  computed: {
    address() {
      return this.$store.state.cosmos.account.address;
    },
    issuer() {
      for (let i = 0; i < this.$store.state.issuers.length; i++) {
        if (this.$store.state.issuers[i].address === this.address) {
          return this.$store.state.issuers[i];
        }
      }
    },
  },
};
</script>
