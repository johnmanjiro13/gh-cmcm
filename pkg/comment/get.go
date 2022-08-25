package comment

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) Get(ctx context.Context, id int64) (*Comment, error) {
	cmt, res, err := c.Repositories.GetComment(ctx, c.owner, c.repo, id)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response: %d", res.StatusCode)
	}

	result := &Comment{}
	if cmt.Body != nil {
		result.Body = *cmt.Body
	}
	if cmt.User != nil && cmt.User.Login != nil {
		result.Author = *cmt.User.Login
	}
	return result, nil
}
