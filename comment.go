package cmcm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v50/github"
)

type commenter struct {
	gh    *github.Client
	owner string
	repo  string
}

type config struct {
	owner string
	repo  string
}

type Comment struct {
	ID      int64  `json:"id"`
	Body    string `json:"body"`
	Author  string `json:"author"`
	HTMLURL string `json:"html_url"`
}

func newCommenter(ctx context.Context, cfg *config) (*commenter, error) {
	cm := &commenter{owner: cfg.owner, repo: cfg.repo}
	cli, err := ghClient(ctx)
	if err != nil {
		return nil, err
	}
	cm.gh = cli
	return cm, nil
}

func (c *commenter) Get(ctx context.Context, id int64) (*Comment, error) {
	comment, _, err := c.gh.Repositories.GetComment(ctx, c.owner, c.repo, id)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	return parseComment(comment), err
}

func (c *commenter) Create(ctx context.Context, sha, body string, path string, position int) (string, error) {
	rc := &github.RepositoryComment{
		Body:     &body,
		Path:     &path,
		Position: &position,
	}
	comment, _, err := c.gh.Repositories.CreateComment(ctx, c.owner, c.repo, sha, rc)
	if err != nil {
		return "", fmt.Errorf("failed to request: %w", err)
	}
	return parseComment(comment).HTMLURL, nil
}

func (c *commenter) Update(ctx context.Context, id int64, body string) (*Comment, error) {
	comment, _, err := c.gh.Repositories.UpdateComment(ctx, c.owner, c.repo, id, &github.RepositoryComment{Body: &body})
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	return parseComment(comment), nil
}

func (c *commenter) List(ctx context.Context, sha string, perPage int) ([]*Comment, error) {
	var page int
	var result []*Comment
	for {
		comments, res, err := c.gh.Repositories.ListCommitComments(ctx, c.owner, c.repo, sha, &github.ListOptions{PerPage: perPage, Page: page})
		if err != nil {
			return nil, fmt.Errorf("failed to request: %w", err)
		}
		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error status code: %d", res.StatusCode)
		}

		for _, comment := range comments {
			result = append(result, parseComment(comment))
		}

		if res.NextPage < 1 {
			break
		}
		page = res.NextPage
	}
	return result, nil
}

func (c *commenter) Delete(ctx context.Context, id int64) error {
	if _, err := c.gh.Repositories.DeleteComment(ctx, c.owner, c.repo, id); err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}
	return nil
}

func parseComment(rc *github.RepositoryComment) *Comment {
	c := &Comment{}
	if rc.ID != nil {
		c.ID = *rc.ID
	}
	if rc.Body != nil {
		c.Body = *rc.Body
	}
	if rc.User != nil && rc.User.Login != nil {
		c.Author = *rc.User.Login
	}
	if rc.HTMLURL != nil {
		c.HTMLURL = *rc.HTMLURL
	}
	return c
}
