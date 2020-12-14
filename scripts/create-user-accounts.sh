#!/bin/sh

# Issuers
EURO_ISSUER_NAME=euro-issuer
EURO_NAME=euro

TOKEN_ISSUER_NAME=token-issuer
TOKEN_NAME=token

CBDC_ISSUER_NAME=cbdc-issuer
CBDC_NAME=cbdc

# Users
USER_1_NAME=euro-user
USER_2_NAME=euro-user-1
USER_3_NAME=euro-user-2


createUser() {
	echo "Creating user key on $1 on node $3\n"

	# create the user key
	docker exec -e USER_NAME=$1 $3 /bin/sh -c 'poacli keys add $(echo $USER_NAME) --keyring-backend test'
	
	# get the address of the user key
	eval ADDRESS=$(docker exec -e USER_NAME=$1 $3 /bin/sh -c 'poacli keys show $(echo $USER_NAME) -a --keyring-backend test')

	# send tokens to the user
	docker exec -e ADDRESS=$ADDRESS $4 /bin/sh -c 'poacli tx send $(poacli keys show validator -a --keyring-backend test) $(echo $ADDRESS) 100000stake -y --trust-node --from validator --chain-id cash --keyring-backend test'

	sleep 5
	
	echo "Creating decentrailsed identifer for user $1\n"
	docker exec -e USER_NAME=$1 $3 /bin/sh -c 'poacli tx did create-did-document --trust-node --from $(echo $USER_NAME) --chain-id cash --keyring-backend test -y'

	sleep 5

	echo "Creating verifiable credential for user $1\n"
	docker exec -e ADDRESS=$ADDRESS $3 /bin/sh -c 'poacli tx did create-verifiable-credential $(echo $ADDRESS) --trust-node --from euro-issuer --chain-id cash --keyring-backend test -y'

	sleep 5

	docker exec -e ADDRESS=$ADDRESS $3 /bin/sh -c 'poacli tx send $(poacli keys show euro-issuer -a --keyring-backend test) $(echo $ADDRESS) 100000euro -y --trust-node --from euro-issuer --chain-id cash --keyring-backend test'

	sleep 5
}

###############################################################################
###                           STEP 1		                            ###
###############################################################################

# Create USERs

echo "Creating user for eurotoken\n"
createUser $USER_1_NAME $EURO_NAME poadnode1 poadnode0 

echo "Creating user for eurotoken\n"
createUser $USER_2_NAME $EURO_NAME poadnode1 poadnode0 

echo "Creating user for eurotoken\n"
createUser $USER_3_NAME $EURO_NAME poadnode1 poadnode0 
