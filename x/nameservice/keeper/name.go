package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/leKg1/nameservice/x/nameservice/types"
)

// CreateName creates a name
func (k Keeper) CreateName(ctx sdk.Context, whois types.Whois) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.WhoisPrefix)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(whois)
	store.Set(key, value)
}

// GetName returns the name information
func (k Keeper) GetName(ctx sdk.Context, whois, key string) (types.Whois, error) {
	store := ctx.KVStore(k.storeKey)
	var name types.Whois
	byteKey := []byte(types.WhoisPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &whois)
	if err != nil {
		return name, err
	}
	return name, nil
}

// DeleteName deletes a name
func (k Keeper) DeleteName(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.WhoisPrefix + key))
}

//
// Functions used by querier
//

// func listName(ctx sdk.Context, k Keeper) ([]byte, error) {
// 	var nameList []types.Name
// 	store := ctx.KVStore(k.storeKey)
// 	iterator := sdk.KVStorePrefixIterator(store, []byte(types.NamePrefix))
// 	for ; iterator.Valid(); iterator.Next() {
// 		var name types.Name
// 		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &name)
// 		nameList = append(nameList, name)
// 	}
// 	res := codec.MustMarshalJSONIndent(k.cdc, nameList)
// 	return res, nil
// }

// func getName(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
// 	key := path[0]
// 	name, err := k.GetName(ctx, key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, err = codec.MarshalJSONIndent(k.cdc, name)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
// 	}

// 	return res, nil
// }

// // Get creator of the item
// func (k Keeper) GetNameOwner(ctx sdk.Context, key string) sdk.AccAddress {
// 	name, err := k.GetName(ctx, key)
// 	if err != nil {
// 		return nil
// 	}
// 	return name.Creator
// }

// // Check if the key exists in the store
// func (k Keeper) NameExists(ctx sdk.Context, key string) bool {
// 	store := ctx.KVStore(k.storeKey)
// 	return store.Has([]byte(types.NamePrefix + key))
// }
