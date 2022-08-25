package comment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/v47/github"
)

func (cli *Client) List(ctx context.Context, sha string, perPage int) (string, error) {
	var page int
	var comments []*Comment
	for {
		cmt, res, err := cli.Repositories.ListCommitComments(ctx, cli.owner, cli.repo, sha, &github.ListOptions{PerPage: perPage, Page: page})
		if err != nil {
			return "", fmt.Errorf("failed to request: %w", err)
		}
		if res.StatusCode != http.StatusOK {
			return "", fmt.Errorf("error response: %d", res.StatusCode)
		}

		for _, c := range cmt {
			cm := &Comment{}
			if c.Body != nil {
				cm.Body = *c.Body
			}
			if c.User != nil && c.User.Login != nil {
				cm.Author = *c.User.Login
			}
			if c.HTMLURL != nil {
				cm.HTMLURL = *c.HTMLURL
			}
			comments = append(comments, cm)
		}

		if res.NextPage < 1 {
			break
		}
		page = res.NextPage
	}
	result, err := json.Marshal(comments)
	if err != nil {
		return "", fmt.Errorf("failed to marshal: %w", err)
	}
	return string(result), nil
}
