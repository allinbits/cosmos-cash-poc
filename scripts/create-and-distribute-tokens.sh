#!/bin/sh

EURO_TOKEN_ISSUER_NAME=euro-token-issuer
EURO_TOKEN_NAME=eurotoken

CBDC_ISSUER_NAME=cbdc-issuer
CBDC_NAME=cbdc

# Creates a issuer with name and token for a given node by a given node
# Takes 4 args (name 0f issuer, name of token, node to create key, node to submit transaction)
createIssuer() {
	echo "Creating issuer key on $1 on node $3\n"

	# create the issuer key
	docker exec -e ISSUER_NAME=$1 $3 /bin/sh -c 'poacli keys add $(echo $ISSUER_NAME) --keyring-backend test'
	
	# get the address of the issuer key
	eval ADDRESS=$(docker exec -e ISSUER_NAME=$1 $3 /bin/sh -c 'poacli keys show $(echo $ISSUER_NAME) -a --keyring-backend test')

	echo $ADDRESS
	echo "Creating issuer with name $1 and token $2: on node $4\n"
	
	## create the issuer token pairing
	docker exec -e ADDRESS=$ADDRESS -e ISSUER_NAME=$1 -e TOKEN_NAME=$2 $4 /bin/sh -c 'poacli tx issuer create-issuer $(echo $ISSUER_NAME) $(echo $ADDRESS) $(echo $TOKEN_NAME) 100000000000 -y --trust-node --from validator --chain-id cash --keyring-backend test'
	
	sleep 5
}

###############################################################################
###                           STEP 1		                            ###
###############################################################################

echo "Creating issuer for eurotoken\n"
createIssuer $EURO_TOKEN_ISSUER_NAME $EURO_TOKEN_NAME poadnode1 poadnode0 

echo "Creating issuer for cdbc\n"
createIssuer $CBDC_ISSUER_NAME $CBDC_NAME poadnode0 poadnode1 


# Mint eurotoken  
docker exec poadnode1 /bin/sh -c "poacli tx issuer mint-token eurotoken 100000000000 -y --trust-node --from euro-token-issuer --chain-id cash --keyring-backend test"

sleep 5

echo "Issuer should have 200000000000 eurotoken's\n"
docker exec poadnode1 /bin/sh -c 'poacli query account $(poacli keys show euro-token-issuer -a --keyring-backend test)'

# Burn euro token
docker exec poadnode1 /bin/sh -c "poacli tx issuer burn-token eurotoken 50000000000 -y --trust-node --from euro-token-issuer --chain-id cash --keyring-backend test"

sleep 5

echo "Issuer should have 150000000000 eurotoken's\n"
docker exec poadnode1 /bin/sh -c 'poacli query account $(poacli keys show euro-token-issuer -a --keyring-backend test)'
