## TODO: set up issuer accounts and distrubte coins

# poadnode1 - create issuer key - echo "y" | go run cmd/poacli/main.go keys add euro-token-issuer
# poadnode2 - create token denom - go run cmd/poacli/main.go tx issuer create-issuer euro-token-issuer $(shell go run cmd/poacli/main.go keys show euro-token-issuer -a) cashmoney 100000000000 --trust-node --from validator --chain-id cash
