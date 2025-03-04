package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"

	"crude/x/crude/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) List(goCtx context.Context, req *types.QueryListRequest) (*types.QueryListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	resources, page, err := k.ListResources(ctx, req.Pagination, req.Name, req.Value)
	if err != nil {
		return nil, errorsmod.Wrap(err, "can't list resources")
	}

	return &types.QueryListResponse{Resources: resources, Pagination: page}, nil
}
