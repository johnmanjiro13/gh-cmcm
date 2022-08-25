package comment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v47/github"
)

func (cli *Client) Update(ctx context.Context, id int64, body string) (string, error) {
	cmt, res, err := cli.Repositories.UpdateComment(ctx, cli.owner, cli.repo, id, &github.RepositoryComment{Body: &body})
	if err != nil {
		return "", fmt.Errorf("failed to request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error response: %d", res.StatusCode)
	}
	var url string
	if cmt.HTMLURL != nil {
		url = *cmt.HTMLURL
	}
	return url, nil
}
