package comment

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

type Client struct {
	*github.Client
	owner string
	repo  string
}

type Config struct {
	Owner string
	Repo  string
}

func NewClient(ctx context.Context, cfg *Config) (*Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, errors.New("github token is missing")
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	baseURL := os.Getenv("GITHUB_BASE_URL")
	if baseURL != "" {
		var err error
		client, err = github.NewEnterpriseClient(baseURL, baseURL, tc)
		if err != nil {
			return nil, fmt.Errorf("failed to create a new github api client: %w", err)
		}
	}

	c := &Client{
		Client: client,
		owner:  cfg.Owner,
		repo:   cfg.Repo,
	}
	return c, nil
}
