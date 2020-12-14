<template>
  <div>
    <div v-for="id in ids" :key="id.ID">
      <md-card-content>
        <div class="md-subtitle">Identity Document</div>
        <md-divider></md-divider>
        <div class="padding">DID: {{ id.id }}</div>
        <div v-for="cred in creds">
          <div class="padding" v-if="cred.id === id.service[0].id">
            Role : {{ cred.credentialsubject.role }}
          </div>
        </div>
      </md-card-content>
    </div>
  </div>
</template>

<style scoped>
.button_spacing {
  margin-top: 5px;
}
.padding {
  padding-top: 10px;
  padding-bottom: 5px;
  padding-left: 10px;
}
</style>

<script>
export default {
  async created() {
    this.$store.dispatch("getCreds");
    this.$store.dispatch("getIds");
  },
  computed: {
    ids() {
      return this.$store.state.ids;
    },
    creds() {
      return this.$store.state.creds;
    },
  },
  methods: {
    async getAll() {
      this.$store.dispatch("getCreds");
      this.$store.dispatch("getIds");
    },
  },
};
</script>
