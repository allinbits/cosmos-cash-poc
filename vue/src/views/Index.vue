<template>
  <div>
    <md-card class="md-elevation-5 sp-container padding">
      <md-card-header>
        <div class="md-layout-item">
          <div class="md-title">Home</div>
          <div class="md-subhead">Check balance and transfer tokens</div>
        </div>
      </md-card-header>
      <div class="md-layout">
        <div class="md-layout-item" />
        <div class="md-layout-item md-size-60">
          <sp-bank-balances />
          <sp-token-send />
        </div>
        <!-- this line is used by starport scaffolding # 4 -->
        <div class="md-layout-item" />
      </div>
    </md-card>
  </div>
</template>

<style scoped>
.padding {
  padding-bottom: 10px;
}
</style>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
