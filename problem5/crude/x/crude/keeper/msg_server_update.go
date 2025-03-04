package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"crude/x/crude/types"
)

func (k msgServer) Update(goCtx context.Context, msg *types.MsgUpdate) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var resource = types.Resource{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Value:   msg.Value,
	}

	_, found, err := k.GetResource(ctx, msg.Id)
	if err != nil {
		return nil, errorsmod.Wrap(err, "can't set resource")
	}
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	err = k.SetResource(ctx, resource)
	if err != nil {
		return nil, errorsmod.Wrap(err, "can't set resource")
	}

	return &types.MsgUpdateResponse{}, nil
}
