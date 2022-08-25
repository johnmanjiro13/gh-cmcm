package comment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v47/github"
)

type CreateOption struct {
	Path     string
	Position int
}

func (cli *Client) Create(ctx context.Context, sha, body string, opt *CreateOption) (string, error) {
	repositoryComment := &github.RepositoryComment{
		Body:     &body,
		Path:     &opt.Path,
		Position: &opt.Position,
	}
	cmt, res, err := cli.Repositories.CreateComment(ctx, cli.owner, cli.repo, sha, repositoryComment)
	if err != nil {
		return "", fmt.Errorf("failed to request: %w", err)
	}
	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("error response: %d", res.StatusCode)
	}

	var url string
	if cmt.HTMLURL != nil {
		url = *cmt.HTMLURL
	}
	return url, nil
}
