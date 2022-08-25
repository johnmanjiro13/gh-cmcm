package cmd_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnmanjiro13/gh-cmcm/pkg/cmd"
	"github.com/johnmanjiro13/gh-cmcm/pkg/comment"
)

func TestPrintPlain(t *testing.T) {
	t.Run("single comment", func(t *testing.T) {
		cmt := &comment.Comment{
			ID:      1,
			Body:    "body",
			Author:  "author",
			HTMLURL: "https://example.com",
		}
		want := `ID:	 1
Author:	 author
URL:	 https://example.com

body
`
		buf := &bytes.Buffer{}
		assert.NoError(t, cmd.PrintPlain(buf, cmt))
		assert.Equal(t, want, buf.String())
	})

	t.Run("multi comment", func(t *testing.T) {
		comments := []*comment.Comment{
			{
				ID:      1,
				Body:    "body",
				Author:  "author",
				HTMLURL: "https://example.com/1",
			},
			{
				ID:      2,
				Body:    "body2",
				Author:  "author2",
				HTMLURL: "https://example.com/2",
			},
		}
		want := `ID:	 1
Author:	 author
URL:	 https://example.com/1

body
----------------------------
ID:	 2
Author:	 author2
URL:	 https://example.com/2

body2
`
		buf := &bytes.Buffer{}
		assert.NoError(t, cmd.PrintPlain(buf, comments...))
		assert.Equal(t, want, buf.String())
	})
}

func TestPrintJSON(t *testing.T) {
	t.Run("single comment", func(t *testing.T) {
		cmt := &comment.Comment{
			ID:      1,
			Body:    "body",
			Author:  "author",
			HTMLURL: "https://example.com",
		}
		want := `{"id":1,"body":"body","author":"author","html_url":"https://example.com"}`
		buf := &bytes.Buffer{}
		assert.NoError(t, cmd.PrintJSON(buf, cmt))
		assert.Equal(t, want, buf.String())
	})

	t.Run("multi comment", func(t *testing.T) {
		comments := []*comment.Comment{
			{
				ID:      1,
				Body:    "body",
				Author:  "author",
				HTMLURL: "https://example.com/1",
			},
			{
				ID:      2,
				Body:    "body2",
				Author:  "author2",
				HTMLURL: "https://example.com/2",
			},
		}
		want := `[{"id":1,"body":"body","author":"author","html_url":"https://example.com/1"},{"id":2,"body":"body2","author":"author2","html_url":"https://example.com/2"}]`
		buf := &bytes.Buffer{}
		assert.NoError(t, cmd.PrintJSON(buf, comments...))
		assert.Equal(t, want, buf.String())
	})
}
