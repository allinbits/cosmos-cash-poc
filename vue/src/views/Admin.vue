<template>
  <div class="sp-container">
    <sp-sign-in />
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

        <md-card-content>
          <br />
          <div class="md-title">Burn Tokens</div>
          <div>
            <md-field>
              <label>Token</label>
              <md-input v-model="token" readonly></md-input>
            </md-field>

            <md-field>
              <label>Burn Token</label>
              <md-input v-model="burn_amount"></md-input>
              <md-icon class="md-accent">warning</md-icon>
            </md-field>

            <div class="md-layout">
              <div class="md-layout-item"></div>
              <md-button
                class="md-raised md-primary button_spacing"
                @click="submit"
              >
                Burn
              </md-button>
            </div>
          </div>
        </md-card-content>
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
  data: () => ({
    token: "cashmoney",
    burn_amount: null,
  }),
  methods: {
    async submit() {
      const payload = {
        type: "issuer/MsgBurnToken",
        body: {
          amount: this.burn_amount,
          token: this.token,
        },
      };
      await this.$store.dispatch("entitySubmit", payload);
    },
  },
};
</script>
