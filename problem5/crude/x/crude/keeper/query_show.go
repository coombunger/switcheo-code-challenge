package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"crude/x/crude/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Show(goCtx context.Context, req *types.QueryShowRequest) (*types.QueryShowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found, err := k.GetResource(ctx, req.Id)
	if err != nil {
		return nil, errorsmod.Wrap(err, "can't get resource")
	}
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", req.Id))
	}

	return &types.QueryShowResponse{Resource: &val}, nil
}
