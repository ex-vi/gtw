package gtw

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

const (
	userFriendshipPath       = "/1.1/friendships/show.json"
	userLookupByUsernamePath = "/2/users/by/username/{username}"
)

type UserFriendshipResponse struct {
	Relationship struct {
		Source struct {
			ID         string `json:"id_str"`
			Username   string `json:"screen_name"`
			Following  bool   `json:"following"`
			FollowedBy bool   `json:"followed_by"`
		} `json:"source"`
		Target struct {
			ID         string `json:"id_str"`
			Username   string `json:"screen_name"`
			Following  bool   `json:"following"`
			FollowedBy bool   `json:"followed_by"`
		} `json:"target"`
	} `json:"relationship"`
}

type UserLookupByUsernameResponse struct {
	Data   User           `json:"data"`
	Errors []PartialError `json:"errors"`
}

// https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/follow-search-get-users/api-reference/get-friendships-show
func (c *Client) UserFriendshipShow(ctx context.Context, sourceUser, targetUser string) (*UserFriendshipResponse, error) {
	p := Parameters{
		QueryParams: map[string]string{
			"source_screen_name": sourceUser,
			"target_screen_name": targetUser,
		},
	}

	res := &UserFriendshipResponse{}
	err := c.Request(ctx, http.MethodGet, userFriendshipPath, p, res)
	if err != nil {
		return nil, errors.Wrap(err, "could not get user friendships")
	}
	return res, nil
}

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func (c *Client) UserLookupByUsername(ctx context.Context, username string) (*UserLookupByUsernameResponse, error) {
	p := Parameters{
		PathParams: map[string]string{
			"username": username,
		},
	}

	res := &UserLookupByUsernameResponse{}
	err := c.Request(ctx, http.MethodGet, userLookupByUsernamePath, p, res)
	if err != nil {
		return nil, errors.Wrap(err, "could not lookup user by username")
	}
	return res, nil
}
