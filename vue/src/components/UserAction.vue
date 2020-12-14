<template>
  <div class="sp-container">
    <md-card-content class="md-layout">
      <br />
      <div class="md-layout-item" />
      <div class="md-layout-item md-size-60">
        <div class="md-title">{{ type }} Tokens</div>
        <md-field>
          <label>Token</label>
          <md-input v-model="token"></md-input>
        </md-field>

        <md-field>
          <label>{{ type }} Token</label>
          <md-input v-model="amount"></md-input>
          <md-icon class="md-accent">warning</md-icon>
        </md-field>

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
  computed: {
    address() {
      return this.$store.state.cosmos.account.address;
    },
  },
  methods: {
    async submit(type) {
      let payload;
      const withdraw = "issuer/MsgWithdrawToken";
      switch (type) {
        case "Withdraw":
          payload = {
            type: withdraw,
            value: {
              token: this.token,
              amount: this.amount,
              owner: this.$store.state.cosmos.account.address,
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
