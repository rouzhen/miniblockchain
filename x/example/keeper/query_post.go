package keeper

import (
	"context"
	"strings"
   "github.com/cosmos/cosmos-sdk/runtime"

	"example/x/example/types"

	"cosmossdk.io/store/prefix"
	//"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//This method retrieves all the posts from the store and returns them as a slice of Post objects.
func (k Keeper) PostAll(ctx context.Context, req *types.QueryAllPostRequest) (*types.QueryAllPostResponse, error) {
	//req validation
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// initialise an empty slice to hold posts retrieved from the store
	var posts []types.Post

	// key-value store from the context using the storeService provided by the Keeper
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Creates a new store with a prefix (namespace) for posts. This ensures only the keys related to posts are queried.
	postStore := prefix.NewStore(store, types.KeyPrefix(types.PostKey))
	// This handles pagination over the postStore. The query.Paginate function iterates over the store entries.
	// The function passed to query.Paginate is called for each key-value pair in the store.
	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var post types.Post
		// Unmarshals the value from the store into the post variable.
		if err := k.cdc.Unmarshal(value, &post); err != nil {
			return err
		}

		// Basic filtering
		//  Checks if a filter string is provided. If the post’s title does not contain the filter string (case-insensitive), it skips adding this post to the result.
		if req.Filter != "" {
			if !strings.Contains(strings.ToLower(post.Title), strings.ToLower(req.Filter)) {
				return nil
			}
		}

		posts = append(posts, post)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Constructs and returns the response with the list of posts and pagination information.
	return &types.QueryAllPostResponse{Post: posts, Pagination: pageRes}, nil
} 

// modified the PostAll method (above) to include a filter parameter



// this method retrieves a post by its id
func (k Keeper) Post(ctx context.Context, req *types.QueryGetPostRequest) (*types.QueryGetPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	post, found := k.GetPost(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPostResponse{Post: post}, nil
}
