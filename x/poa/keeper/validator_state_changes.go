package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ApplyAndReturnValidatorSetUpdates at the end of every block we update and return the validator set
func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate) {
	validators := k.GetAllValidators(ctx)

	for _, validator := range validators {
		if validator.Accepted == true {
			updates = append(updates, validator.ABCIValidatorUpdate())
		}
	}

	return updates
}

// CalculateValidatorVote happens at the start of every block to ensure no malacious actors
func (k Keeper) CalculateValidatorVotes(ctx sdk.Context) {
	// TODO: On every block calculate and update validator set

	validators := k.GetAllValidators(ctx)

	// TODO: Smart query method
	for _, validator := range validators {
		votes := k.GetAllVotesForValidator(ctx, (validator.Name))
		if len(votes) == len(validators) {
			validator.Accepted = true
			k.SetValidator(ctx, validator.Name, validator)
		}
	}

	//votes := k.GetAllVotes(ctx)

	// TODO: Brute force method
	//var count int = 0
	//for _, validator := range validators {
	//	for _, vote := range votes {
	//		if vote.Name == validator.Name {
	//			count++
	//			if count == len(validators) {
	//			}
	//		}
	//	}
	//}

	// TODO: Jail validators if malicious

}
