#!/bin/sh

###############################################################################
###                           STEP 1		                            ###
###############################################################################

echo "Creating issuer for eurotoken\n"

# Create euro-token-issuer and eurotoken
docker exec poadnode1 /bin/sh -c "poacli keys add euro-token-issuer --keyring-backend test"
eval ADDRESS=$(docker exec poadnode1 /bin/sh -c "poacli keys show euro-token-issuer -a --keyring-backend test")
docker exec -e ADDRESS=$ADDRESS poadnode0 /bin/sh -c 'poacli tx issuer create-issuer euro-token-issuer $(echo $ADDRESS) eurotoken 100000000000 -y --trust-node --from validator --chain-id cash --keyring-backend test'

sleep 5

echo "Issuer should have 100000000000 eurotoken's\n"
docker exec poadnode1 /bin/sh -c 'poacli query account $(poacli keys show euro-token-issuer -a --keyring-backend test)'

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
