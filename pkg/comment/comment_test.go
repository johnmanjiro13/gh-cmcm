package comment_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/johnmanjiro13/gh-cmcm/pkg/comment"
	"github.com/johnmanjiro13/gh-cmcm/pkg/comment/mock_comment"
)

func TestCommenter_Get(t *testing.T) {
	const id int64 = 1
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := map[string]struct {
		cmt *comment.Comment
		err error
	}{
		"success":   {&comment.Comment{Body: "test body"}, nil},
		"has error": {nil, errors.New("an error")},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := mock_comment.NewMockClient(ctrl)
			client.EXPECT().GetComment(gomock.Any(), id).Return(tt.cmt, tt.err)
			commenter := comment.NewCommenter(client)
			got, err := commenter.Get(context.Background(), id)

			assert.Equal(t, tt.cmt, got)
			if tt.err != nil {
				assert.Error(t, tt.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCommenter_Create(t *testing.T) {
	const (
		sha  = "sha"
		body = "body"
	)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := map[string]struct {
		cmt *comment.Comment
		err error
	}{
		"success":   {&comment.Comment{HTMLURL: "https://example.com"}, nil},
		"has error": {nil, errors.New("an error")},
	}

	opt := &comment.CreateOption{}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := mock_comment.NewMockClient(ctrl)
			client.EXPECT().CreateComment(gomock.Any(), sha, body, opt).Return(tt.cmt, tt.err)
			commenter := comment.NewCommenter(client)
			got, err := commenter.Create(context.Background(), sha, body, opt)

			if tt.err != nil {
				assert.Equal(t, "", got)
				assert.Error(t, tt.err, err)
			} else {
				assert.Equal(t, tt.cmt.HTMLURL, got)
				assert.NoError(t, err)
			}
		})
	}
}

func TestCommenter_Update(t *testing.T) {
	const (
		id   int64 = 1
		body       = "body"
	)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := map[string]struct {
		cmt *comment.Comment
		err error
	}{
		"success":   {&comment.Comment{HTMLURL: "https://example.com"}, nil},
		"has error": {nil, errors.New("an error")},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := mock_comment.NewMockClient(ctrl)
			client.EXPECT().UpdateComment(gomock.Any(), id, body).Return(tt.cmt, tt.err)
			commenter := comment.NewCommenter(client)
			got, err := commenter.Update(context.Background(), id, body)

			if tt.err != nil {
				assert.Equal(t, "", got)
				assert.Error(t, tt.err, err)
			} else {
				assert.Equal(t, tt.cmt.HTMLURL, got)
				assert.NoError(t, err)
			}
		})
	}
}

func TestCommenter_List(t *testing.T) {
	const (
		sha     = "sha"
		perPage = 30
	)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := map[string]struct {
		comments []*comment.Comment
		err      error
	}{
		"success":   {[]*comment.Comment{{Body: "body"}}, nil},
		"has error": {nil, errors.New("an error")},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := mock_comment.NewMockClient(ctrl)
			client.EXPECT().ListComment(gomock.Any(), sha, perPage).Return(tt.comments, tt.err)
			commenter := comment.NewCommenter(client)
			got, err := commenter.List(context.Background(), sha, perPage)

			assert.Equal(t, tt.comments, got)
			if tt.err != nil {
				assert.Error(t, tt.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCommenter_Delete(t *testing.T) {
	const id int64 = 1
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := map[string]struct {
		err error
	}{
		"success":   {nil},
		"has error": {errors.New("an error")},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := mock_comment.NewMockClient(ctrl)
			client.EXPECT().DeleteComment(gomock.Any(), id).Return(tt.err)
			commenter := comment.NewCommenter(client)
			err := commenter.Delete(context.Background(), id)

			if tt.err != nil {
				assert.Error(t, tt.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
