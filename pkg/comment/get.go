package comment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (cli *Client) Get(ctx context.Context, id int64) (string, error) {
	cmt, res, err := cli.Repositories.GetComment(ctx, cli.owner, cli.repo, id)
	if err != nil {
		return "", fmt.Errorf("failed to request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error response: %d", res.StatusCode)
	}

	c := &Comment{}
	if cmt.Body != nil {
		c.Body = *cmt.Body
	}
	if cmt.User != nil && cmt.User.Login != nil {
		c.Author = *cmt.User.Login
	}
	if cmt.HTMLURL != nil {
		c.HTMLURL = *cmt.HTMLURL
	}
	result, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("failed to marshal: %w", err)
	}
	return string(result), nil
}
