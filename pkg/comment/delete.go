package comment

import (
	"context"
	"fmt"
	"net/http"
)

func (cli *Client) Delete(ctx context.Context, id int64) error {
	res, err := cli.Repositories.DeleteComment(ctx, cli.owner, cli.repo, id)
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}
	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("error response: %d", res.StatusCode)
	}
	return nil
}
