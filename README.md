### Cosmos-Cash

---

Cosmos cash is a distributed application that stores data in a [key/value store](https://www.techopedia.com/definition/26284/key-value-store) and secures that data using a [Proof Of Authority](https://changelly.com/blog/what-is-proof-of-authority-poa/) consensus algorithm.

The goal of the application is to re-define how an [electric money institution](https://thebanks.eu/emis) works by leveraging the [cosmos-sdk](https://github.com/cosmos/cosmos-sdk/) and [tendermint](https://github.com/tendermint/tendermint/).

### How to install the application

---

1. Clone the repository 

```sh
git clone git@github.com:allinbits/cosmos-cash-poa.git 
```

<br />

2. Install the binaries (poad, poacli)

```
cd cosmos-cash-poa
make install
```

<br />

3. Ensure binaries are available

```
poad -h
poacli -h
```

<br />

### How to initialize the application 

---

1. Initialize the genesis file ($HOME/.poad/config/genesis.json)

```sh
poad init --chain-id=cash cash
```

<br />

2. Create a key for the first validator

```sh
poacli keys add validator
```
<br />

3. Add the validator that was created in the step 2 as the first validator and assign them 1000000000 `cash` coins

```sh
poad add-genesis-account $(poacli keys show validator -a) 1000000000cash,1000000000stake

```

<br />

4. Generate a initial `CreateValidator` transaction to allow other applications in the network to sync when they join

```sh
poad gentx --name validator
```

<br />

5. Put the previously generated transaction in the correct location to allow the application to start correctly

```sh
poad collect-gentxs
```

<br />

6. Start the applicaton :tada:

```sh
poad start
```

<br />

### How to use the application

---

#### How to add a validator to the validator set

1. Create a CreateValidatorPOA transaction and submit it to the applicaton

```sh
poacli tx poa create-validator val1 `hex-encoded-public-key-of-the-validator-you-want-to-vote-for` --trust-node --from validator --chain-id cash
```

<br />

2. Verfiy that the CreateValidatorPOA tranaction was correctly processed by the application

```sh
poacli query poa validator-poa val1 --trust-node --chain-id cash
```

<br />

#### How to vote on a validator to allow the validator to be added to the validator set (consensus)

1. Vote for a validator to join the validator set (consensus)

```sh
poacli tx poa vote-validator val1 --trust-node --from validator --chain-id cash
```

<br />

2. Query for the vote to see if the transaction was successful

```sh
poacli query poa vote-poa val1 `validator address` --trust-node --chain-id cash
```

<br />

### How to find more information on the application

---

Please view to Proof of Authority module spec for more details on the commands defined in the `How to` sections above

> **[Link to Spec](./x/poa/spec/README.md)**

