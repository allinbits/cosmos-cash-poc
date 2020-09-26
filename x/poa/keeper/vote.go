package keeper

import (
	"github.com/allinbits/cosmos-cash-poa/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVote sets a vote with key as votee and voter combined in a byte array
func (k Keeper) SetVote(ctx sdk.Context, vote types.Vote) {
	k.Set(ctx, append([]byte(vote.Name), vote.Voter...), types.VotesKey, vote)

	k.Set(ctx, append(vote.Voter, []byte(vote.Name)...), types.VotesByValidatorKey, vote)
}

func (k Keeper) GetVote(ctx sdk.Context, key []byte) (types.Vote, bool) {
	vote, found := k.Get(ctx, key, types.VotesKey, k.UnmarshalVote)
	return vote.(types.Vote), found
}

func (k Keeper) DeleteVote(ctx sdk.Context, key []byte) {
	k.Delete(ctx, key, types.VotesKey)
}

func (k Keeper) UnmarshalVote(value []byte) (interface{}, bool) {
	vote := types.Vote{}
	err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &vote)
	if err != nil {
		return types.Vote{}, false
	}
	return vote, true
}

func (k Keeper) GetAllVotes(ctx sdk.Context) (votes []types.Vote) {
	val := k.GetAll(ctx, types.VotesKey, k.UnmarshalVote)

	// TODO: Make this nicer
	for _, value := range val {
		votes = append(votes, value.(types.Vote))
	}

	return votes
}

func (k Keeper) GetAllVotesForValidator(ctx sdk.Context, name string) (votes []types.Vote) {
	val := k.GetAll(ctx, append(types.VotesKey, []byte(name)...), k.UnmarshalVote)

	for _, value := range val {
		vote := value.(types.Vote)
		if vote.InFavor == true {
			votes = append(votes, vote)
		}
	}

	return votes
}

func (k Keeper) PurgeAllVotesByValidator(ctx sdk.Context, voter sdk.ValAddress) bool {
	val := k.GetAll(ctx, append(types.VotesByValidatorKey, voter...), k.UnmarshalVote)

	for _, value := range val {
		vote := value.(types.Vote)
		k.DeleteVote(ctx, append([]byte(vote.Name), vote.Voter...))
	}

	k.Delete(ctx, voter, types.VotesByValidatorKey)

	return true
}
