package service

import (
	"context"
	"errors"

	"github.com/wojciechpawlinow/find-indexes/internal/domain/index"
)

type IndexFinderPort interface {
	Find(ctx context.Context, v int) (int, bool, error)
}

type IndexFinderService struct {
	Repo index.Repository
}

func (s *IndexFinderService) Find(ctx context.Context, v int) (int, bool, error) {

	if v < 0 {
		return 0, false, errors.New("invalid range")
	}

	direct := true

	i, err := s.Repo.OneByValue(ctx, v)
	if err != nil {
		direct = false
		i, err = s.Repo.OneByValue(ctx, v-(v*10/100))
		if err != nil {
			i, err = s.Repo.OneByValue(ctx, v+(v*10/100))
		}
	}

	return i, direct, err
}
