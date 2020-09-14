package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ApplyAndReturnValidatorSetUpdates at the end of every block we update and return the validator set
func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate) {
	validators := k.GetAllAcceptedValidators(ctx)

	for _, validator := range validators {
		updates = append(updates, validator.ABCIValidatorUpdate())
	}

	return updates
}

// CalculateValidatorVote happens at the start of every block to ensure no malacious actors
func (k Keeper) CalculateValidatorVotes(ctx sdk.Context) {
	acceptedValidators := k.GetAllAcceptedValidators(ctx)
	validators := k.GetAllValidators(ctx)
	qourum := k.GetParams(ctx).Quorum

	// Smart query method
	for _, validator := range validators {
		votes := k.GetAllVotesForValidator(ctx, (validator.Name))

		// check the number of votes are greater that the qourum needed
		if float32(len(votes)) >= (float32(len(acceptedValidators)))*(float32(qourum)/100) || len(validators) == 1 {
			validator.Accepted = true
			k.SetValidator(ctx, validator.Name, validator)
		}
	}

	// TODO: Brute force method
	//votes := k.GetAllVotes(ctx)
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
