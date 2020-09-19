<template>
  <div>
    <div class="button_box">
      <button
        :class="[
          'button',
        ]"
        @click="getIssuers"
      >
        Get Issuers
      </button>
    </div>
    <div>
       <h3>ISSUER LIST</h3>
    </div>
    <div class="issuer_list" v-for="issuer in issuers" >
      <div :key="issuer.address">
        <p>Name: {{issuer.name}}</p>
        <p>Address: {{issuer.address}}</p>
        <button
          :class="[
            'button',
          ]"
          @click="getTokens(issuer.address)"
	>
	  Token: {{issuer.token}}
	</button>
        <div class="token_box" v-for="token in tokens" >
          <div :key="token.denom">
            <p>Name: {{token.denom}}</p>
            <p>Address: {{token.amount}}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

.issuer_list {
  box-shadow: inset 0 0 0 1px rgba(0, 0, 0, 0.1);
  margin-bottom: 1rem;
  padding: 1rem;
  border-radius: 0.5rem;
  overflow: hidden;
  align-items: center;
  justify-content: center;
  font-size: 0.85rem;
  letter-spacing: 0.05em;
  border-radius: 0.25rem;
  display: flex;
}
.token_box {
  margin-bottom: 1rem;
  padding: 1rem;
  border-radius: 0.5rem;
  overflow: hidden;
  align-items: center;
  justify-content: center;
  font-size: 0.85rem;
  letter-spacing: 0.05em;
  display: flex;
}
.button_box {
  padding-top: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
}
button {
  background: none;
  border: none;
  color: rgba(0, 125, 255);
  padding: 0;
  font-size: inherit;
  font-weight: 800;
  font-family: inherit;
  text-transform: uppercase;
  margin-top: 0.5rem;
  cursor: pointer;
  transition: opacity 0.1s;
  letter-spacing: 0.03em;
  transition: color 0.25s;
  display: inline-flex;
  align-items: center;
}
</style>

<script>

export default {
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
        this.$store.dispatch("getIssuers", { val: "val" })
      },
    async getTokens(address) {
        this.$store.dispatch("getTokens", { address })
      }
    },
  };

</script>
