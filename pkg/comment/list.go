package comment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/v47/github"
)

type Comment struct {
	ID      int64  `json:"id"`
	Body    string `json:"body"`
	Author  string `json:"author"`
	HTMLURL string `json:"html_url"`
}

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
			comments = append(comments, parseComment(c))
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

func parseComment(cmt *github.RepositoryComment) *Comment {
	c := &Comment{}
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
