#!/bin/sh

# TODO: finish off setting up network 
# TODO: this is temporary should use 
# TODO: use docker exec to run commands
# TODO: use a integration test

### STEP 1

# NODE0
# Enter the docker machine
docker exec -it poanode0 /bin/sh

# Import the exported key
poacli keys import validator validator

# Create the validator
poacli tx poa create-validator validator $(poad tendermint show-validator) --trust-node --from validator --chain-id cash

# Have the validator vote for itself
poacli tx poa vote-validator validator --trust-node --from validator --chain-id cash

### STEP 2

# NODE1
# Create the key for the node
docker exec -it poanode1 /bin/sh
poacli keys add validator1
poacli key show validator1 -a

# NODE2
# Create the key for the node
docker exec -it poanode2 /bin/sh
poacli keys add validator2
poacli keys show validator2 -a

# NODE3
# Create the key for the node
docker exec -it poanode3 /bin/sh
poacli keys add validator3
poacli keys show validator3 -a

### STEP 3

# NODE0
# Send tokens to each validator 
poacli tx send $NODE0ADDRESS $NODE1ADDRESS 10000stake --from validator --trust-node --chain-id cash
poacli tx send $NODE0ADDRESS $NODE2ADDRESS 10000stake --from validator --trust-node --chain-id cash
poacli tx send $NODE0ADDRESS $NODE3ADDRESS 10000stake --from validator --trust-node --chain-id cash

### STEP 4

# NODE1
# Create validator for validator set and vote for it
poacli tx poa create-validator validator1 $(poad tendermint show-validator) --trust-node --from validator1 --chain-id cash
poacli tx poa vote-validator validator1 --trust-node --from validator1 --chain-id cash

# NODE0
# Vote for validator1 to join the set
poacli tx poa vote-validator validator1 --trust-node --from validator --chain-id cash

### STEP 5

# Verify valdiators are in the set by checking the proposer address of the block
curl 0.0.0.0:26657/block?height?803 | jq '.result.block.header.proposer_address'
