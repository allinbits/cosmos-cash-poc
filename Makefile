PACKAGES=$(shell go list ./...)

#VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
#COMMIT := $(shell git log -1 --format='%H')

# TODO: Update the ldflags with the app, client & server names
#ldflags = -X github.com/allinbits/cosmos-cash-poa/version.Name=POC \
#	-X github.com/allinbits/cosmos-cash-poa/version.ServerName=poad \
#	-X github.com/allinbits/cosmos-cash-poa/version.ClientName=poacli \
#	-X github.com/allinbits/cosmos-cash-poa/version.Version=$(VERSION) \
#	-X github.com/allinbits/cosmos-cash-poa/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

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

# TODO: Remove at a later date when POC is at closing stage

init-dev: init-chain init-validator

init-chain:
	go run cmd/poad/main.go cmd/poad/genaccounts.go init --chain-id=cash cash
	go run cmd/poacli/main.go keys add validator

init-validator:
	go run cmd/poad/main.go cmd/poad/genaccounts.go add-genesis-account $(shell go run cmd/poacli/main.go keys show validator -a) 1000000000stake
	go run cmd/poad/main.go cmd/poad/genaccounts.go gentx --name validator
	go run cmd/poad/main.go cmd/poad/genaccounts.go collect-gentxs

create-validator:
	go run cmd/poacli/main.go tx poa create-validator validator $(shell go run cmd/poad/main.go cmd/poad/genaccounts.go tendermint show-validator) --trust-node --from validator --chain-id cash

query-validator:
	go run cmd/poacli/main.go query poa validator-poa validator --trust-node --chain-id cash

query-all-validators:
	go run cmd/poacli/main.go query poa validators

vote-validator:
	go run cmd/poacli/main.go tx poa vote-validator validator --trust-node --from validator --chain-id cash

query-vote:
	go run cmd/poacli/main.go query poa vote-poa validator $(shell go run cmd/poacli/main.go keys show validator --bech val -a) --trust-node --chain-id cash

query-all-votes:
	go run cmd/poacli/main.go query poa votes


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
