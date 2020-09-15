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

init-chain:
	go run cmd/poad/main.go cmd/poad/genaccounts.go init --chain-id=cash cash --home ./build/.poad
	echo "y" | go run cmd/poacli/main.go keys add validator

init-validator:
	go run cmd/poad/main.go cmd/poad/genaccounts.go add-genesis-account $(shell go run cmd/poacli/main.go keys show validator -a) 1000000000stake --home ./build/.poad
	go run cmd/poad/main.go cmd/poad/genaccounts.go gentx --name validator --home ./build/.poad
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

localnet-set-up-consensus:
	./scripts/set-up-poa-network.sh


localnet-start: init-dev export-key
	NODE0ADDRESS=$(shell go run cmd/poad/main.go cmd/poad/genaccounts.go tendermint show-node-id --home ./build/.poad)@192.16.10.2:26656 docker-compose up

export-key:
	echo "password1234\npassword1234" | poacli keys export validator 2> ./build/validator

###############################################################################
###                           Helpful Commands	                            ###
###############################################################################

# TODO: Remove at a later date when POC is at closing stage

create-validator:
	go run cmd/poacli/main.go tx poa create-validator validator $(shell go run cmd/poad/main.go cmd/poad/genaccounts.go tendermint show-validator) --trust-node --from validator --chain-id cash --home ./build/.poad

query-validator:
	go run cmd/poacli/main.go query poa validator-poa validator --trust-node --chain-id cash --home ./build/.poad

query-all-validators:
	go run cmd/poacli/main.go query poa validators --home ./build/.poad

vote-validator:
	go run cmd/poacli/main.go tx poa vote-validator validator --trust-node --from validator --chain-id cash --home ./build/.poad

query-vote:
	go run cmd/poacli/main.go query poa vote-poa validator $(shell go run cmd/poacli/main.go keys show validator --bech val -a) --trust-node --chain-id cash --home ./build/.poad

query-all-votes:
	go run cmd/poacli/main.go query poa votes --home ./build/.poad

send-coin:
	poacli tx send cosmos1h6c36qvkpdycas468lcfmp6xjl39cwvethemuj cosmos1djh6qxf893a6lk2evecce5cexhvan5tcpnezt0 10000stake --from validator --trust-node --chain-id cash

.PHONY:				\
	test			\
	lint			\
	init-dev 		\
	init-chain		\
	create-validator	\
	query-validator		\
	query-all-validators	\
	vote-validator		\
	query-vote		\
	query-add-votes		\
	send-coin		\
	clean			\
	export-key		\
	build-linux		\
	localnet-start		\
