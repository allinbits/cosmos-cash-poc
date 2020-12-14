#!/bin/sh


# Issuers
EURO_USER_NAME=euro-issuer
EURO_NAME=euro

TOKEN_USER_NAME=token-issuer
TOKEN_NAME=token

CBDC_USER_NAME=cbdc-issuer
CBDC_NAME=cbdc

# Users
USER_1_NAME=euro-user


createUser() {
	echo "Creating issuer key on $1 on node $3\n"

	# create the issuer key
	docker exec -e USER_NAME=$1 $3 /bin/sh -c 'poacli keys add $(echo $USER_NAME) --keyring-backend test'
	
	# get the address of the issuer key
	eval ADDRESS=$(docker exec -e USER_NAME=$1 $3 /bin/sh -c 'poacli keys show $(echo $USER_NAME) -a --keyring-backend test')

	# send tokens to the issuer
	docker exec -e ADDRESS=$ADDRESS $4 /bin/sh -c 'poacli tx send $(poacli keys show validator -a --keyring-backend test) $(echo $ADDRESS) 100000stake -y --trust-node --from validator --chain-id cash --keyring-backend test'

	sleep 5
	
	echo "Creating decentrailsed identifer for issuer $1\n"
	docker exec -e USER_NAME=$1 $3 /bin/sh -c 'poacli tx did create-did-document --trust-node --from $(echo $USER_NAME) --chain-id cash --keyring-backend test -y'

	sleep 5

	echo "Creating verifiable credential for issuer $1\n"
	docker exec -e ADDRESS=$ADDRESS $4 /bin/sh -c 'poacli tx did create-verifiable-credential $(echo $ADDRESS) --trust-node --from euro-issuer --chain-id cash --keyring-backend test -y'

	sleep 5
}
