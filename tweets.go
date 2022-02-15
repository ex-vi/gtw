package gtw

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

const (
	tweetRetweetedBy = "/2/tweets/{tweetID}/retweeted_by"
)

type TweetRetweetedByResponse struct {
	Data   []User         `json:"data"`
	Meta   PaginationMeta `json:"meta,omitempty"`
	Errors []PartialError `json:"errors,omitempty"`
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/quick-start/retweets-lookup
func (c *Client) TweetRetweetedByUsers(ctx context.Context, tweetID string) ([]User, error) {
	users := []User{}

	nextToken := ""
	for {
		res, err := c.tweetRetweetedBy(ctx, tweetID, nextToken)
		if err != nil {
			return nil, errors.Wrap(err, "could not get tweet retweeters")
		}
		users = append(users, res.Data...)

		nextToken = StringValue(res.Meta.NextToken)
		if nextToken == "" {
			break
		}
	}
	return users, nil
}

func (c *Client) tweetRetweetedBy(ctx context.Context, tweetID, nextToken string) (*TweetRetweetedByResponse, error) {
	p := Parameters{
		PathParams: map[string]string{
			"tweetID": tweetID,
		},
	}

	if nextToken != "" {
		p.QueryParams = map[string]string{
			"pagination_token": nextToken,
		}
	}

	res := &TweetRetweetedByResponse{}
	err := c.Request(ctx, http.MethodGet, tweetRetweetedBy, p, res)
	if err != nil {
		return nil, errors.Wrap(err, "could not get tweet retweeters page")
	}
	return res, nil
}
