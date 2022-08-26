package api_test

import (
	"testing"

	"github.com/google/go-github/v47/github"
	"github.com/stretchr/testify/assert"

	"github.com/johnmanjiro13/gh-cmcm/pkg/api"
	"github.com/johnmanjiro13/gh-cmcm/pkg/comment"
)

func TestParseComment(t *testing.T) {
	tests := map[string]struct {
		cmt  *github.RepositoryComment
		want *comment.Comment
	}{
		"has id":         {&github.RepositoryComment{ID: toPointer(int64(1))}, &comment.Comment{ID: 1}},
		"has body":       {&github.RepositoryComment{Body: toPointer("body")}, &comment.Comment{Body: "body"}},
		"has user login": {&github.RepositoryComment{User: &github.User{Login: toPointer("login")}}, &comment.Comment{Author: "login"}},
		"has html url":   {&github.RepositoryComment{HTMLURL: toPointer("https://example.com")}, &comment.Comment{HTMLURL: "https://example.com"}},
		"has user only":  {&github.RepositoryComment{User: &github.User{}}, &comment.Comment{}},
		"blank":          {&github.RepositoryComment{}, &comment.Comment{}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, api.ParseComment(tt.cmt))
		})
	}
}

func toPointer[T any](i T) *T {
	return &i
}
