package cmcm

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func ghClient(ctx context.Context) (*github.Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, errors.New("github token is missing. please use GITHUB_TOKEN environment variable")
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
	return client, nil
}
