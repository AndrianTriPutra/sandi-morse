package morse

import "context"

type repo struct {
}

func NewRepository() RepositoryI {
	return &repo{}
}

type RepositoryI interface {
	Decode(ctx context.Context, req string) []string
	Encode(ctx context.Context, req string) []string
}
