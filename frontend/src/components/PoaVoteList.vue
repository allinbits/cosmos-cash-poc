<template>
  <div>
    <div class="button_box">
      <button
        :class="[
          'button',
        ]"
        @click="getPoaVotes"
      >
        Get Poa Votes
      </button>
    </div>
    <div>
       <h3>VOTES LIST</h3>
    </div>
    <div class="vote_list" v-for="vote in votes">
      <div :key="vote.voter">
        <p>Name: {{vote.name}}</p>
        <p>Voter: {{vote.voter}}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>

.vote_list {
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
.button_box {
  padding-top: 50px;
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
    votes() {
      return this.$store.state.poaVotes;
    },
  },
  methods: {
    async getPoaVotes() {
        this.$store.dispatch("getPoaVotes", { val: "val" })
      }
    }
  };

</script>
