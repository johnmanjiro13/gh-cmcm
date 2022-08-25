package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"

	"github.com/johnmanjiro13/gh-cmcm/pkg/comment"
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

	c := &Client{
		Client: client,
		owner:  cfg.Owner,
		repo:   cfg.Repo,
	}
	return c, nil
}

func (cli *Client) GetComment(ctx context.Context, id int64) (*comment.Comment, error) {
	cmt, res, err := cli.Repositories.GetComment(ctx, cli.owner, cli.repo, id)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response: %d", res.StatusCode)
	}

	return parseComment(cmt), nil
}

func (cli *Client) CreateComment(ctx context.Context, sha, body string, opt *comment.CreateOption) (*comment.Comment, error) {
	rc := &github.RepositoryComment{
		Body:     &body,
		Path:     &opt.Path,
		Position: &opt.Position,
	}
	cmt, res, err := cli.Repositories.CreateComment(ctx, cli.owner, cli.repo, sha, rc)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("error response: %d", res.StatusCode)
	}

	return parseComment(cmt), nil
}

func (cli *Client) UpdateComment(ctx context.Context, id int64, body string) (*comment.Comment, error) {
	cmt, res, err := cli.Repositories.UpdateComment(ctx, cli.owner, cli.repo, id, &github.RepositoryComment{Body: &body})
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response: %d", res.StatusCode)
	}
	return parseComment(cmt), nil
}

func (cli *Client) ListComment(ctx context.Context, sha string, perPage int) ([]*comment.Comment, error) {
	var page int
	var comments []*comment.Comment
	for {
		cmt, res, err := cli.Repositories.ListCommitComments(ctx, cli.owner, cli.repo, sha, &github.ListOptions{PerPage: perPage, Page: page})
		if err != nil {
			return nil, fmt.Errorf("failed to request: %w", err)
		}
		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error response: %d", res.StatusCode)
		}

		for _, c := range cmt {
			comments = append(comments, parseComment(c))
		}

		if res.NextPage < 1 {
			break
		}
		page = res.NextPage
	}

	return comments, nil
}

func (cli *Client) DeleteComment(ctx context.Context, id int64) error {
	res, err := cli.Repositories.DeleteComment(ctx, cli.owner, cli.repo, id)
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}
	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("error response: %d", res.StatusCode)
	}
	return nil
}

func parseComment(cmt *github.RepositoryComment) *comment.Comment {
	c := &comment.Comment{}
	if cmt.ID != nil {
		c.ID = *cmt.ID
	}
	if cmt.Body != nil {
		c.Body = *cmt.Body
	}
	if cmt.User != nil && cmt.User.Login != nil {
		c.Author = *cmt.User.Login
	}
	if cmt.HTMLURL != nil {
		c.HTMLURL = *cmt.HTMLURL
	}
	return c
}
