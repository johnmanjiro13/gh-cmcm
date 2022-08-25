package comment

import (
	"context"
	"fmt"

	"github.com/google/go-github/v47/github"
)

func (c *Client) List(ctx context.Context, sha string) ([]*Comment, error) {
	var page int
	var comments []*Comment
	for {
		cmt, res, err := c.Repositories.ListCommitComments(ctx, c.owner, c.repo, sha, &github.ListOptions{PerPage: 1, Page: page})
		if err != nil {
			return nil, fmt.Errorf("failed to request: %w", err)
		}

		for _, c := range cmt {
			cm := &Comment{}
			if c.Body != nil {
				cm.Body = *c.Body
			}
			if c.User != nil && c.User.Login != nil {
				cm.Author = *c.User.Login
			}
			comments = append(comments, cm)
		}

		if res.NextPage < 1 {
			break
		}
		page = res.NextPage
	}
	return comments, nil
}
