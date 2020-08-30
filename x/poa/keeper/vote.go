package keeper

import (
	"github.com/allinbits/cosmos-cash-poa/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVote sets a vote with key as votee and voter combined in a byte array
func (k Keeper) SetVote(ctx sdk.Context, vote types.Vote) {
	k.Set(ctx, append([]byte(vote.Name), vote.Voter...), types.VotesKey, vote)
}

func (k Keeper) GetVote(ctx sdk.Context, key []byte) (types.Vote, bool) {
	vote, found := k.Get(ctx, key, types.VotesKey, k.UnmarshalVote)
	return vote.(types.Vote), found
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

	// TODO: Make this nicer
	for _, value := range val {
		votes = append(votes, value.(types.Vote))
	}

	return votes
}
