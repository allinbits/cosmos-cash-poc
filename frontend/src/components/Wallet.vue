<template>
  <div>
    <md-card>
      <md-card-header class="md-layout">
      <div class="md-layout-item">
        <div class="md-title">{{ address ? "Your Account" : "Sign in" }}</div>
        <div class="md-subhead">View your tokens</div>
      </div>
      </md-card-header>

      <div v-if="!address" class="password">
        <input
          type="text"
          v-model="password"
          class="password__input"
          placeholder="Password (mnemonic)"
        />
        <div
          :class="[
            'button',
            `button__error__${!!error}`,
            `button__enabled__${!!mnemonicValid}`,
          ]"
          @click="signIn"
        >
          Sign in
        </div>
      </div>
      <div v-else class="account endPadding">
        <div class="card">
          <div class="card__row">
            <div class="md-primary card__icon">
	      <md-icon class="icon_spacing md-primary">person</md-icon>
            </div>
              {{ address }}
          </div>
          <div v-if="account.coins">
            <span class="md-primary"> 
		 <md-list>
		      <md-list-item>
        		<md-icon class="md-primary"></md-icon>
        		<span class="md-list-item-text md-subhead">Token Name</span>
        		<span class="md-list-item-text md-subhead">Amount in Pool</span>
      			</md-list-item>
			<md-divider></md-divider>
              <span
                v-for="coin in account.coins"
                :key="coin.denom"
                >
		      <md-list-item :key="coin.denom">
        		<md-icon>request_page</md-icon>
        		<span class="md-list-item-text">{{coin.denom}}</span>
        		<span class="md-list-item-text md-primary">{{coin.amount}}</span>
      			</md-list-item>
			<md-divider></md-divider>
	</span>
		 </md-list>
            </span>
          </div>
        </div>

    </div>

      </md-card>
  </div>
</template>

<style scoped>
.container {
  margin-bottom: 1.5rem;
}
.card {
  border-radius: 0.25rem;
  padding: 0.25rem 0.75rem;
  overflow-x: hidden;
  margin-right: 40px;
  margin-left: 40px;
  padding-bottom: 20px;
}
.card__row {
  display: flex;
  align-items: center;
  margin: 0.5rem 0;
  color: rgba(0, 0, 0, 0.25);
  font-size: 0.875rem;
  font-weight: 500;
  line-height: 1.5;
}
.card__icon {
  width: 1.75rem;
  height: 1.75rem;
  fill: rgba(0, 0, 0, 0.15);
  flex-shrink: 0;
}
.card__desc {
  letter-spacing: 0.02em;
  padding: 0 0.5rem;
  word-break: break-all;
}
.h1 {
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 1rem;
}
.password {
  margin-top: 0.5rem;
  display: flex;
  margin-left: 40px;
    margin-right: 40px;
    padding-bottom: 30px;
}
.password__input {
  border: none;
  width: 100%;
  padding: 0.75rem;
  box-sizing: border-box;
  font-family: inherit;
  background: rgba(0, 0, 0, 0.03);
  font-size: 0.85rem;
  border-radius: 0.25rem;
  color: rgba(0, 0, 0, 0.5);
}
.button {
  margin-left: 1rem;
  background: rgba(0, 0, 0, 0.03);
  padding: 0 1.5rem;
  white-space: nowrap;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 0.85rem;
  color: rgba(0, 0, 0, 0.25);
  text-transform: uppercase;
  letter-spacing: 0.05em;
  border-radius: 0.25rem;
  transition: all 0.1s;
  user-select: none;
}
.button.button__error__true {
  animation: shake 0.82s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
  background: rgba(255, 228, 228, 0.5);
  color: rgb(255, 0, 0);
}
.button__enabled__false {
  cursor: not-allowed;
}
.button__enabled__true {
  color: rgba(0, 125, 255);
  font-weight: 700;
  cursor: pointer;
}
.button__enabled__true:active {
  color: rgba(0, 125, 255, 0.65);
}
.password__input:focus {
  outline: none;
  background: rgba(0, 0, 0, 0.06);
  color: rgba(0, 0, 0, 0.5);
}
.password__input::placeholder {
  color: rgba(0, 0, 0, 0.35);
  font-weight: 500;
}
.coin__amount {
  text-transform: uppercase;
  font-size: 0.75rem;
  letter-spacing: 0.02em;
  font-weight: 600;
}
.coin__amount:after {
  content: ",";
  margin-right: 0.25em;
}
.coin__amount:last-child:after {
  content: "";
  margin-right: initial;
}
@keyframes shake {
  10%,
  90% {
    transform: translate3d(-1px, 0, 0);
  }
  20%,
  80% {
    transform: translate3d(2px, 0, 0);
  }
  30%,
  50%,
  70% {
    transform: translate3d(-4px, 0, 0);
  }
  40%,
  60% {
    transform: translate3d(4px, 0, 0);
  }
}
.endPadding{
padding-bottom: 20px;

}
</style>

<script>
import * as bip39 from "bip39";

export default {
  data() {
    return {
      password: "",
      error: false
    };
  },
  computed: {
    account() {
      return this.$store.state.account;
    },
    address() {
      const { client } = this.$store.state;
      const address = client && client.senderAddress;
      return address;
    },
    mnemonicValid() {
      return bip39.validateMnemonic(this.passwordClean);
    },
    passwordClean() {
      return this.password.trim();
    }
  },
  methods: {
    async signIn() {
      if (this.mnemonicValid && !this.error) {
        const mnemonic = this.passwordClean;
        this.$store.dispatch("accountSignIn", { mnemonic }).catch(() => {
          this.error = true;
          setTimeout(() => {
            this.error = false;
          }, 1000);
        });
      }
    }
  }
};
</script>
