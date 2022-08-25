.PHONY: test mockgen

test:
	go test -cover -race ./...

mockgen:
	mockgen -destination pkg/comment/mock_comment/mock_comment.go -package mock_comment github.com/johnmanjiro13/gh-cmcm/pkg/comment Client