package keeper

import (
	"context"

	"crude/x/crude/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Create(goCtx context.Context, msg *types.MsgCreate) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var resource = types.Resource{
		Creator: msg.Creator,
		Name:    msg.Name,
		Value:   msg.Value,
	}

	id := k.AppendResource(ctx, resource)

	return &types.MsgCreateResponse{Id: id}, nil
}
