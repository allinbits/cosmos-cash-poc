<template>
  <div>
    <md-card class="md-elevation-5">
      <md-card-header class="md-layout">
        <div class="md-layout-item">
          <div class="md-title">Validator List</div>
          <div class="md-subhead">All validators in the application</div>
        </div>
        <md-button
          class="md-raised md-primary button_spacing"
          v-on:click="getPoaValidators"
          ><md-icon>refresh</md-icon></md-button
        >
      </md-card-header>

      <md-card-content>
        <br />
        <md-table>
          <md-table-row>
            <md-table-head>Moniker</md-table-head>
            <md-table-head>Address</md-table-head>
            <md-table-head>In Validator Set</md-table-head>
            <!--- <md-table-head>Has Been Approved</md-table-head> --->
            <md-table-head>Votes</md-table-head>
          </md-table-row>
          <md-table-row
            v-for="validator in validators"
            :key="validator.address"
          >
            <md-table-cell>{{ validator.description.moniker }}</md-table-cell>
            <md-table-cell class="address">{{
              validator.address
            }}</md-table-cell>
            <md-table-cell v-if="validator.inset">
              <md-icon class="icon_spacing md-primary">done_outline</md-icon>
            </md-table-cell>
            <md-table-cell v-else>
              <md-icon class="icon_spacing">highlight_off</md-icon>
            </md-table-cell>
            <!--- <md-table-cell v-if="validator.accepted">
	        <md-icon class="md-primary icon_spacing" >done_outline</md-icon>
	      </md-table-cell>
              <md-table-cell v-else>
	        <md-icon class="icon_spacing">highlight_off</md-icon>
	      </md-table-cell>
	      --->
            <md-table-cell
              ><md-menu :md-offset-x="127" :md-offset-y="-36">
                <md-button md-menu-trigger>
                  <md-icon>how_to_vote</md-icon>
                </md-button>
                <md-menu-content>
                  <div v-for="vote in votes">
                    <div v-if="vote.name === validator.name">
                      <div v-if="vote.in_favor">
                        <md-menu-item class="icon_spacing_popup">
                          <md-icon class="md-primary">done_outline</md-icon>
                        </md-menu-item>
                      </div>
                      <div v-else>
                        <md-menu-item class="icon_spacing_popup">
                          <md-icon>highlight_off</md-icon>
                        </md-menu-item>
                      </div>
                    </div>
                  </div>
                </md-menu-content>
              </md-menu>
            </md-table-cell>
          </md-table-row>
        </md-table>
      </md-card-content>
    </md-card>
  </div>
</template>
<style scoped>
.icon_spacing {
  margin-left: 30px;
}
.icon_spacing_popup {
  margin-left: 10px;
}
.button_spacing {
  margin-right: 30px;
  margin-top: 5px;
}
.address {
  overflow: hidden;
  max-width: 350px;
}
</style>
<script>
export default {
  name: "TableBasic",
  async created() {
    this.$store.dispatch("getPoaValidators", { val: "val" });
    this.$store.dispatch("getPoaVotes", { val: "val" });
  },
  computed: {
    validators() {
      return this.$store.state.poaValidators;
    },
    votes() {
      return this.$store.state.poaVotes;
    },
  },
  methods: {
    async getPoaValidators() {
      this.$store.dispatch("getPoaValidators", { val: "val" });
      this.$store.dispatch("getPoaVotes", { val: "val" });
    },
    async getPoaVotes() {
      this.$store.dispatch("getPoaVotes", { val: "val" });
    },
  },
  isValidatorVote(validator) {
    console.log(validator);
  },
};
</script>
