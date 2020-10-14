<template>
  <div class="sp-container">
    <div v-if="token">
      <md-card-content class="md-layout">
        <br />
        <div class="md-layout-item" />
        <div class="md-layout-item md-size-60">
          <div class="md-title">{{ type }} Tokens</div>
          <md-field>
            <label>Token</label>
            <md-input v-model="token" readonly></md-input>
          </md-field>

          <div v-if="type !== 'Freeze' && type !== 'Unfreeze'">
            <md-field>
              <label>{{ type }} Token</label>
              <md-input v-model="amount"></md-input>
              <md-icon class="md-accent">warning</md-icon>
            </md-field>
          </div>

          <div class="md-layout">
            <div class="md-layout-item"></div>
            <div v-if="loading">
              <md-progress-spinner
                :md-diameter="30"
                md-mode="indeterminate"
              ></md-progress-spinner>
            </div>
            <md-button
              class="md-raised md-primary button_spacing"
              @click="submit(type)"
              :disabled="loading"
            >
              {{ type }} Token
            </md-button>
          </div>
        </div>
        <div class="md-layout-item" />
      </md-card-content>
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
  props: {
    type: String,
  },
  data: () => ({
    amount: null,
    loading: false,
  }),
  async beforeCreate() {
    this.$store.dispatch("getIssuers");
  },
  computed: {
    token() {
      if (this.$store.state.tokens[this.$store.state.cosmos.account.address]) {
        return this.$store.state.tokens[
          this.$store.state.cosmos.account.address
        ][0].denom;
      }
    },
    address() {
      return this.$store.state.cosmos.account.address;
    },
  },
  methods: {
    async submit(type) {
      let payload;
      const burn = "issuer/MsgBurnToken";
      const mint = "issuer/MsgMintToken";
      const freeze = "issuer/MsgFreezeToken";
      const unfreeze = "issuer/MsgUnfreezeToken";
      switch (type) {
        case "Burn":
          payload = {
            type: burn,
            value: {
              amount: this.amount,
              token: this.token,
              issuer: this.$store.state.cosmos.account.address,
            },
          };

          break;
        case "Mint":
          payload = {
            type: mint,
            value: {
              amount: this.amount,
              token: this.token,
              issuer: this.$store.state.cosmos.account.address,
            },
          };
          break;
        case "Freeze":
          payload = {
            type: freeze,
            value: {
              token: this.token,
              issuer: this.$store.state.cosmos.account.address,
            },
          };
          break;
        case "Unfreeze":
          payload = {
            type: unfreeze,
            value: {
              token: this.token,
              issuer: this.$store.state.cosmos.account.address,
            },
          };
          break;
      }

      this.loading = true;
      await this.$store.dispatch("entitySubmit", payload);
      this.loading = false;
    },
  },
};
</script>
