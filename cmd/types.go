package cmd

import "context"

type Comment struct {
	Body   string
	Author string
}

type Commenter interface {
	List(ctx context.Context, sha string) ([]*Comment, error)
}
