#!/bin/sh

###############################################################################
###                           FUNCTIONS		                            ###
###############################################################################

# Creates a validator for a given node
# Take 1 arg the name of the node e.g poadnode0
createValidator() {
	echo "Creating validator for node $1\n"

	# Create the validator
	docker exec $1 /bin/sh -c 'poacli tx poa create-validator $(poacli keys show validator --bech val -a --keyring-backend test) $(poad tendermint show-validator) moniker identity website security@contact details -y --trust-node --from validator --chain-id cash --keyring-backend test'

	sleep 5
}

# Votes for a perspecitve canidate
# Take 2 args the name of the node voting and the candidate node e.g poadnode0 poadnode1
voteForValidator() {
	eval CANDIDATE=$(docker exec $2 /bin/sh -c "poacli keys show validator --bech val -a --keyring-backend test")
	echo "Votee $1 is voting for candidate $2"
	docker exec -e CANDIDATE=$CANDIDATE $1 /bin/sh -c 'poacli tx poa vote-validator $(echo $CANDIDATE) -y --trust-node --from validator --chain-id cash --keyring-backend test'

	sleep 5
}

# Kicks for a perspecitve canidate
# Take 2 args the name of the node voting and the candidate node e.g poadnode0 poadnode1
kickValidator() {
	eval CANDIDATE=$(docker exec $2 /bin/sh -c "poacli keys show validator --bech val -a --keyring-backend test")
	echo "Votee $1 is voting to kick candidate $2"
	docker exec -e CANDIDATE=$CANDIDATE $1 /bin/sh -c 'poacli tx poa kick-validator $(echo $CANDIDATE) -y --trust-node --from validator --chain-id cash --keyring-backend test'

	sleep 5
}
###############################################################################
###                           STEP 1		                            ###
###############################################################################

# Import the exported key for the first node
docker exec poadnode0 /bin/sh -c "echo -e 'password1234\n' | poacli keys import validator validator --keyring-backend test"

## Create the validator
voteForValidator poadnode0 poadnode0

###############################################################################
###                           STEP 2		                            ###
###############################################################################

# Create the keys for each node
for var in poadnode1 poadnode2 poadnode3
do
	echo "Creating key for node $var\n"
	docker exec $var /bin/sh -c "poacli keys add validator --keyring-backend test"
done


## Send tokens to each validator
for node in poadnode1 poadnode2 poadnode3
do
	eval ADDRESS=$(docker exec $node /bin/sh -c "poacli keys show validator -a --keyring-backend test")
	echo "Sending tokens to $ADDRESS\n"
	docker exec -e ADDRESS=$ADDRESS poadnode0 /bin/sh -c 'poacli tx send $(poacli keys show validator -a --keyring-backend test) $(echo $ADDRESS) 100000stake -y --trust-node --from validator --chain-id cash --keyring-backend test'
	sleep 5
done

###############################################################################
###                           STEP 3		                            ###
###############################################################################

# Create validator for validator set
for var in poadnode1 poadnode2 poadnode3
do
	createValidator $var
done

###############################################################################
###                           STEP 4		                            ###
###############################################################################

# Adding new validators to the set

# Vote for validator1 to join the set
voteForValidator poadnode0 poadnode1

# poadnode1 votes for poadnode0 to prove the node is in the consensus
voteForValidator poadnode1 poadnode0

# poadnode1 votes for poadnode1 to stay relevant in the consensus
voteForValidator poadnode1 poadnode1

# poadnode1 and poanode0 votes for poadnode2 to join the consensus
voteForValidator poadnode0 poadnode2
voteForValidator poadnode1 poadnode2

# poadnode2 votes for poadnode2 to stay relevant in the consensus
voteForValidator poadnode2 poadnode2

# poadnode2 votes for poadnode1 to prove the node is in the consensus
voteForValidator poadnode2 poadnode1

# poadnode2 votes for poadnode0 to prove the node is in the consensus
voteForValidator poadnode2 poadnode0

# kick poadnode2 out of the consensus
kickValidator poadnode0 poadnode2
kickValidator poadnode1 poadnode2

echo "POA Consensus started with 2 nodes :thumbs_up:\n"

## Verify valdiators are in the set by checking the proposer address of the block
#curl 0.0.0.0:26657/block?height?803 | jq '.result.block.header.proposer_address'
## Verify valdiators are in the set by checking the validator set
#curl -X GET "localhost/validators?height=50&page=1&per_page=30" -H  "accept: application/json"
