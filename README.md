### Cosmos-Cash

---

Cosmos cash is a distributed application that stores data in a [key/value store](https://www.techopedia.com/definition/26284/key-value-store) and is secured using a [Proof Of Authority](https://github.com/allinbits/modules/tree/master/x/poa) consensus algorithm.

The goal of the application is to re-define how an [electric money institution](https://thebanks.eu/emis) works by leveraging the [cosmos-sdk](https://github.com/cosmos/cosmos-sdk/) and [tendermint](https://github.com/tendermint/tendermint/).

### How to install the application

---

1. Clone the repository

```sh
git clone git@github.com:allinbits/cosmos-cash-poc.git
```

<br />

2. Install the binaries (poad, poacli)

```
cd cosmos-cash-poc
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

7. Run commands outlined in the Makefile

```sh
cat Makefile | grep create-
```

<br />

### How to run a localnet

---

1. Build the dockerfile

```sh
docker build -t 388991194029.dkr.ecr.us-east-1.amazonaws.com/allinbits-dev/cosmos-cash-poa .
```

<br />

2. Run the localnet

```sh
make localnet-start
```

<br />

3. Set up the consensus

```sh
make localnet-consensus
```

<br />

4. Create issuer data

```sh
make localnet-distribute-tokens
```

<br />

4. Create user data

```sh
make localnet-users
```

<br />

### How to run the webui (run localnet commands for seed data)

---

1. Go to the `vue` folder

```sh
cd vue
```

<br />

2. Start the web server

```sh
yarn serve
```

<br />

3. Check the website in the browser

[localhost:8080](http://localhost:8080)

<br />

