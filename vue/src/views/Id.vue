<template>
  <div>
    <md-card class="md-elevation-5 sp-container padding">
      <md-card-header>
        <div class="md-layout-item">
          <div class="md-title">IDs</div>
          <div class="md-subhead">
            All did documents and verifiable credentials
          </div>
        </div>
      </md-card-header>
      <div class="md-layout">
        <div class="md-layout-item" />
        <div class="md-layout-item md-size-60">
          <id-list />
        </div>
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
