PACKAGES=$(shell go list ./...)

# TODO: Update the ldflags with the app

#VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
#COMMIT := $(shell git log -1 --format='%H')

#ldflags = -X github.com/allinbits/cosmos-cash-poa/version.Name=POC \
#	-X github.com/allinbits/cosmos-cash-poa/version.ServerName=poad \
#	-X github.com/allinbits/cosmos-cash-poa/version.ClientName=poacli \
#	-X github.com/allinbits/cosmos-cash-poa/version.Version=$(VERSION) \
#	-X github.com/allinbits/cosmos-cash-poa/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

###############################################################################
###                           Basic Golang Commands                         ###
###############################################################################

all: install

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/poad
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/poacli

install-debug: go.sum
	go build -mod=readonly $(BUILD_FLAGS) -gcflags="all=-N -l" ./cmd/poad
	go build -mod=readonly $(BUILD_FLAGS) -gcflags="all=-N -l" ./cmd/poacli

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)

lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

###############################################################################
###                           Chain Initialization                          ###
###############################################################################

init-dev: init-chain init-validator
	
start-dev:
	go run cmd/poad/main.go cmd/poad/genaccounts.go start --home ./build/.poad

init-chain:
	go run cmd/poad/main.go cmd/poad/genaccounts.go init --chain-id=cash cash --home ./build/.poad
	echo "y" | go run cmd/poacli/main.go keys add validator

init-validator:
	go run cmd/poad/main.go cmd/poad/genaccounts.go add-genesis-account $(shell go run cmd/poacli/main.go keys show validator -a) 1000000000stake --home ./build/.poad
	go run cmd/poad/main.go cmd/poad/genaccounts.go gentx --name validator --home ./build/.poad --moniker cash --website test.com --identity test --security-contact test@test.com --details atest
	go run cmd/poad/main.go cmd/poad/genaccounts.go collect-gentxs --home ./build/.poad

clean:
	sudo rm -r ./build
	docker-compose down

build-linux:
	mkdir -p ./build
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./build ./cmd/poad
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./build ./cmd/poacli

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

localnet-distribute-tokens:
	./scripts/create-and-distribute-tokens.sh

localnet-consensus:
	./scripts/set-up-poa-network.sh


localnet-start: init-dev export-key
	NODE0ADDRESS=$(shell go run cmd/poad/main.go cmd/poad/genaccounts.go tendermint show-node-id --home ./build/.poad)@192.16.10.2:26656 docker-compose up

export-key:
	echo "password1234\npassword1234" | poacli keys export validator 2> ./build/validator

###############################################################################
###                           Helpful Commands	                            ###
###############################################################################

# TODO: Remove at a later date when POC is at closing stage

### poa module commands

create-validator:
	go run cmd/poacli/main.go tx poa create-validator $(shell go run cmd/poacli/main.go keys show validator --bech val -a) $(shell go run cmd/poad/main.go cmd/poad/genaccounts.go tendermint show-validator) moniker identity website securityContact details --trust-node --from validator --chain-id cash --home ./build/.poad

query-validator:
	go run cmd/poacli/main.go query poa validator-poa $(shell go run cmd/poacli/main.go keys show validator --bech val -a) --trust-node --chain-id cash --home ./build/.poad

query-all-validators:
	go run cmd/poacli/main.go query poa validators --home ./build/.poad

vote-validator:
	go run cmd/poacli/main.go tx poa vote-validator $(shell go run cmd/poacli/main.go keys show validator --bech val -a) --trust-node --from validator --chain-id cash --home ./build/.poad

kick-validator:
	go run cmd/poacli/main.go tx poa kick-validator $(shell go run cmd/poacli/main.go keys show validator --bech val -a) --trust-node --from validator --chain-id cash --home ./build/.poad

query-vote:
	go run cmd/poacli/main.go query poa vote-poa $(shell go run cmd/poacli/main.go keys show validator --bech val -a) $(shell go run cmd/poacli/main.go keys show validator --bech val -a) --trust-node --chain-id cash --home ./build/.poad

query-all-votes:
	go run cmd/poacli/main.go query poa votes --home ./build/.poad


### issuer module commands

create-custom-issuer-key:
	echo "y" | go run cmd/poacli/main.go keys add $$CUSTOM_TOKEN-issuer

create-custom-issuer: create-custom-issuer-key
	go run cmd/poacli/main.go tx issuer create-issuer $$CUSTOM_TOKEN-issuer $(shell go run cmd/poacli/main.go keys show $$CUSTOM_TOKEN-issuer -a) $$CUSTOM_TOKEN 100000000000 --trust-node --from validator --chain-id cash --home ./build/.poad

create-issuer-key:
	echo "y" | go run cmd/poacli/main.go keys add euro-token-issuer

create-issuer: create-issuer-key
	go run cmd/poacli/main.go tx issuer create-issuer euro-token-issuer $(shell go run cmd/poacli/main.go keys show euro-token-issuer -a) eurotoken 100000000000 --trust-node --from validator --chain-id cash --home ./build/.poad

query-all-issuers:
	go run cmd/poacli/main.go query issuer issuers --home ./build/.poad

send-token:
	poacli tx send $(shell go run cmd/poacli/main.go keys show euro-token-issuer -a) $(shell go run cmd/poacli/main.go keys show validator -a) 50000000eurotoken --from euro-token-issuer -y --trust-node --chain-id cash

send-token-back:
	poacli tx send $(shell go run cmd/poacli/main.go keys show validator -a) $(shell go run cmd/poacli/main.go keys show euro-token-issuer -a) 50000000eurotoken --from euro-token-issuer -y --trust-node --chain-id cash

mint-token: 
	go run cmd/poacli/main.go tx issuer mint-token eurotoken 50000000 --trust-node --from euro-token-issuer --chain-id cash

burn-token: 
	go run cmd/poacli/main.go tx issuer burn-token eurotoken 50000000 --trust-node --from euro-token-issuer --chain-id cash

withdraw-token: 
	go run cmd/poacli/main.go tx issuer withdraw-token eurotoken 50000000 --trust-node --from euro-token-issuer --chain-id cash

freeze-token: 
	go run cmd/poacli/main.go tx issuer freeze-token eurotoken --trust-node --from euro-token-issuer --chain-id cash

unfreeze-token: 
	go run cmd/poacli/main.go tx issuer unfreeze-token eurotoken --trust-node --from euro-token-issuer --chain-id cash

query-balance: 
	go run cmd/poacli/main.go query account $(shell go run cmd/poacli/main.go keys show validator -a)

.PHONY:					\
	test				\
	lint				\
	init-dev 			\
	init-chain			\
	create-validator		\
	query-validator			\
	query-all-validators		\
	vote-validator			\
	query-vote			\
	query-add-votes			\
	send-coin			\
	clean				\
	export-key			\
	build-linux			\
	localnet-start			\
	localnet-set-up-consensus	\
	create-issuer-key		\
	create-issuer			\
	query-balance			\
