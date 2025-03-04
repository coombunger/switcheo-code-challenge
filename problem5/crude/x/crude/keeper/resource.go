package keeper

import (
	"cosmossdk.io/store/prefix"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"crude/x/crude/types"
)

func (k Keeper) AppendResource(ctx sdk.Context, resource types.Resource) uint64 {
	count := k.GetResourceCount(ctx)
	resource.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))
	appendedValue := k.cdc.MustMarshal(&resource)
	store.Set(GetResourceIDBytes(resource.Id), appendedValue)
	k.SetResourceCount(ctx, count+1)
	return count
}

func (k Keeper) GetResourceCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ResourceCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetResourceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetResourceCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ResourceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetResource(ctx sdk.Context, id uint64) (val types.Resource, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))
	b := store.Get(GetResourceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetResource(ctx sdk.Context, resource types.Resource) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))
	b := k.cdc.MustMarshal(&resource)
	store.Set(GetResourceIDBytes(resource.Id), b)
}

func (k Keeper) DeleteResource(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))
	store.Delete(GetResourceIDBytes(id))
}

func (k Keeper) ListResources(ctx sdk.Context, pageReq *query.PageRequest, nameFilter string, valueFilter string) ([]types.Resource, *query.PageResponse, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))

	iterator := store.Iterator(pageReq.Key, nil)
	defer iterator.Close()

	offset := pageReq.Offset
	filteredCount := uint64(0)
	totalCount := uint64(0)
	if pageReq.Key != nil {
		offset = 0
		totalCount = binary.BigEndian.Uint64(pageReq.Key)
	}

	var resources []types.Resource
	for ; iterator.Valid(); iterator.Next() {
		var resource types.Resource

		err := k.cdc.Unmarshal(iterator.Value(), &resource)
		if err != nil {
			return nil, nil, err
		}

		if (nameFilter == "" || resource.Name == nameFilter) &&
			(valueFilter == "" || resource.Value == valueFilter) {
			if filteredCount >= offset+pageReq.Limit {
				break
			}

			if filteredCount >= offset {
				resources = append(resources, resource)
			}
			filteredCount++
		}
		totalCount += 1
	}

	pageRes := &query.PageResponse{NextKey: nil}

	if filteredCount >= offset+pageReq.Limit {
		bz := make([]byte, 8)
		binary.BigEndian.PutUint64(bz, totalCount)
		pageRes.NextKey = bz
	}

	return resources, pageRes, nil
}
