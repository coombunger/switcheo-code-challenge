package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"crude/x/crude/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Delete(goCtx context.Context, msg *types.MsgDelete) (*types.MsgDeleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found, err := k.GetResource(ctx, msg.Id)
	if err != nil {
		return nil, errorsmod.Wrap(err, "can't delete resource")
	}
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	err = k.DeleteResource(ctx, msg.Id)
	if err != nil {
		return nil, errorsmod.Wrap(err, "can't delete resource")
	}

	return &types.MsgDeleteResponse{}, nil
}
