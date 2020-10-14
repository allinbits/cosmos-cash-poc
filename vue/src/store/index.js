import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import { coins } from "@cosmjs/launchpad";

import cosmos from "@tendermint/vue/src/store/cosmos.js";

const VUE_APP_API_COSMOS = "http://127.0.0.1:1317";

Vue.use(Vuex);
export default new Vuex.Store({
  modules: { cosmos },
  state: {
    validators: [],
    poaValidators: [],
    poaVotes: [],
    issuers: [],
    tokens: [],
  },
  mutations: {
    entitySet(state, { type, body }) {
      const updated = {};
      updated[type] = body;
      state.data = { ...state.data, ...updated };
    },
    clientUpdate(state, { client }) {
      state.client = client;
    },
    validatorsUpdate(state, validators) {
      state.validators = validators;
    },
    poaValidatorsUpdate(state, poaValidators) {
      state.poaValidators = poaValidators;
    },
    poaVotesUpdate(state, poaVotes) {
      state.poaVotes = poaVotes;
    },
    issuersUpdate(state, issuers) {
      state.issuers = issuers;
    },
    tokensUpdate(state, tokens) {
      state.tokens = tokens;
    },
  },
  actions: {
    async entitySubmit({ state }, msg) {
      const { chain_id } = state.cosmos;
      const creator = state.cosmos.client.senderAddress;
      const memo = "admin action";
      const fee = {
        amount: coins(200, msg.value.token),
        gas: "200000",
      };

      return await state.cosmos.client.signAndPost([msg], fee, memo);
    },
    async getValidators({ state, commit }, { type, body }) {
      const { data } = await axios.get(`${TENDERMINT}/validators`);
      commit("validatorsUpdate", data.result.validators);
    },
    async getPoaValidators({ state, commit }) {
      const { data } = await axios.get(`${VUE_APP_API_COSMOS}/poa/validators`);
      commit("poaValidatorsUpdate", data.result);
    },
    async getPoaVotes({ state, commit }) {
      const { data } = await axios.get(`${VUE_APP_API_COSMOS}/poa/votes`);
      commit("poaVotesUpdate", data.result);
    },
    async getIssuers({ state, commit }) {
      const { data } = await axios.get(`${VUE_APP_API_COSMOS}/issuer/issuers`);
      data.result.map(async (issuer) => {
        await this.dispatch("getTokens", { address: issuer.address });
      });
      commit("issuersUpdate", data.result);
    },
    async getTokens({ state, commit }, { address }) {
      const { data } = await axios.get(
        `${VUE_APP_API_COSMOS}/bank/balances/${address}`
      );

      state.tokens[address] = data.result;
      commit("tokensUpdate", state.tokens);
    },
  },
});
