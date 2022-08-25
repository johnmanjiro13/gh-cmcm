package comment

import (
	"context"
	"encoding/json"
	"fmt"
)

type Client interface {
	GetComment(ctx context.Context, id int64) (*Comment, error)
	CreateComment(ctx context.Context, sha, body string, opt *CreateOption) (*Comment, error)
	UpdateComment(ctx context.Context, id int64, body string) (*Comment, error)
	ListComment(ctx context.Context, sha string, perPage int) ([]*Comment, error)
	DeleteComment(ctx context.Context, id int64) error
}

type Comment struct {
	ID      int64  `json:"id"`
	Body    string `json:"body"`
	Author  string `json:"author"`
	HTMLURL string `json:"html_url"`
}

type Commenter struct {
	client Client
}

func NewCommenter(client Client) *Commenter {
	return &Commenter{client: client}
}

type CreateOption struct {
	Path     string
	Position int
}

func (c *Commenter) Get(ctx context.Context, id int64) (string, error) {
	cmt, err := c.client.GetComment(ctx, id)
	if err != nil {
		return "", err
	}
	result, err := json.Marshal(cmt)
	if err != nil {
		return "", fmt.Errorf("failed to marshal: %w", err)
	}
	return string(result), nil
}

func (c *Commenter) Create(ctx context.Context, sha, body string, opt *CreateOption) (string, error) {
	cmt, err := c.client.CreateComment(ctx, sha, body, opt)
	if err != nil {
		return "", err
	}
	return cmt.HTMLURL, err
}

func (c *Commenter) Update(ctx context.Context, id int64, body string) (string, error) {
	cmt, err := c.client.UpdateComment(ctx, id, body)
	if err != nil {
		return "", err
	}
	return cmt.HTMLURL, nil
}

func (c *Commenter) List(ctx context.Context, sha string, perPage int) (string, error) {
	comments, err := c.client.ListComment(ctx, sha, perPage)
	if err != nil {
		return "", err
	}
	result, err := json.Marshal(comments)
	if err != nil {
		return "", fmt.Errorf("failed to marshal: %w", err)
	}
	return string(result), nil
}

func (c *Commenter) Delete(ctx context.Context, id int64) error {
	return c.client.DeleteComment(ctx, id)
}
