package comment

import (
	"context"
)

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=mock_$GOPACKAGE/mock_$GOPACKAGE.go
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

func (c *Commenter) Get(ctx context.Context, id int64) (*Comment, error) {
	cmt, err := c.client.GetComment(ctx, id)
	if err != nil {
		return nil, err
	}
	return cmt, nil
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

func (c *Commenter) List(ctx context.Context, sha string, perPage int) ([]*Comment, error) {
	comments, err := c.client.ListComment(ctx, sha, perPage)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *Commenter) Delete(ctx context.Context, id int64) error {
	return c.client.DeleteComment(ctx, id)
}
