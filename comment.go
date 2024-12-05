package cmcm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/tomnomnom/linkheader"
)

type commenter struct {
	client *api.RESTClient
	owner  string
	repo   string
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

func newCommenter(cfg *config) (*commenter, error) {
	cm := &commenter{owner: cfg.owner, repo: cfg.repo}
	client, err := api.DefaultRESTClient()
	if err != nil {
		return nil, fmt.Errorf("faled to create rest client: %w", err)
	}
	cm.client = client
	return cm, nil
}

func (c *commenter) Get(id int64) (*Comment, error) {
	resp := struct {
		ID       int64
		Body     string
		User     struct{ Login string }
		HTML_URL string
	}{}
	if err := c.client.Get(fmt.Sprintf("repos/%s/%s/comments/%d", c.owner, c.repo, id), &resp); err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	return &Comment{
		ID:      resp.ID,
		Body:    resp.Body,
		Author:  resp.User.Login,
		HTMLURL: resp.HTML_URL,
	}, nil
}

func (c *commenter) Create(sha, body string, path string, position int) (string, error) {
	reqBody := struct {
		Body     string `json:"body"`
		Path     string `json:"path"`
		Position int    `json:"position"`
	}{
		Body:     body,
		Path:     path,
		Position: position,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(reqBody); err != nil {
		return "", fmt.Errorf("failed to encode request body: %w", err)
	}
	resp := struct{ HTML_URL string }{}
	if err := c.client.Post(fmt.Sprintf("repos/%s/%s/commits/%s/comments", c.owner, c.repo, sha), &buf, &resp); err != nil {
		return "", fmt.Errorf("failed to request: %w", err)
	}
	return resp.HTML_URL, nil
}

func (c *commenter) Update(id int64, body string) (string, error) {
	reqBody := struct {
		Body string `json:"body"`
	}{Body: body}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(reqBody); err != nil {
		return "", fmt.Errorf("failed to encode request body: %w", err)
	}
	resp := struct{ HTML_URL string }{}
	if err := c.client.Patch(fmt.Sprintf("repos/%s/%s/comments/%d", c.owner, c.repo, id), &buf, &resp); err != nil {
		return "", fmt.Errorf("failed to request: %w", err)
	}
	return resp.HTML_URL, nil
}

func (c *commenter) List(sha string, perPage int) ([]*Comment, error) {
	var result []*Comment

	path := fmt.Sprintf("repos/%s/%s/commits/%s/comments?per_page=%d", c.owner, c.repo, sha, perPage)
	for {
		resp, err := c.client.Request(http.MethodGet, path, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to request: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error status code: %d", resp.StatusCode)
		}

		comments := make([]struct {
			ID       int64
			Body     string
			User     struct{ Login string }
			HTML_URL string
		}, 0)
		if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
			return nil, fmt.Errorf("failed to decode response body: %w", err)
		}
		for _, c := range comments {
			result = append(result, &Comment{
				ID:      c.ID,
				Body:    c.Body,
				Author:  c.User.Login,
				HTMLURL: c.HTML_URL,
			})
		}
		links := linkheader.Parse(resp.Header.Get("Link"))
		if len(links) == 0 {
			break
		}

		var hasNext bool
		for _, link := range links {
			if link.Rel == "next" {
				hasNext = true
				u, err := url.Parse(link.URL)
				if err != nil {
					return nil, fmt.Errorf("failed to parse url: %w", err)
				}
				path = fmt.Sprintf("%s?%s", u.Path[1:], u.Query().Encode())
			}
		}
		if !hasNext {
			break
		}
	}
	return result, nil
}

func (c *commenter) Delete(id int64) error {
	if err := c.client.Delete(fmt.Sprintf("repos/%s/%s/comments/%d", c.owner, c.repo, id), nil); err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}
	return nil
}
