<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <issuer-list class="top_spacing" />
      <!-- this line is used by starport scaffolding # 4 -->
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
};
</script>
