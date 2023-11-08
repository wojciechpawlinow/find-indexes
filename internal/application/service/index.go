package service

import (
	"context"
	"errors"

	"github.com/wojciechpawlinow/find-indexes/internal/domain/index"
)

type IndexPort interface {
	FindByValue(ctx context.Context, v int) (int, bool, error)
}

type IndexService struct {
	Repo index.Repository
}

var _ IndexPort = (*IndexService)(nil)

// FindByValue matches the exact value in the storage or the closest one
func (s *IndexService) FindByValue(ctx context.Context, v int) (int, bool, error) {

	if v < 0 {
		return 0, false, errors.New("invalid value range")
	}

	directMatch := true

	i, err := s.Repo.OneByValue(ctx, v)
	if err != nil {
		directMatch = false
		i, err = s.Repo.OneByValue(ctx, v-(v*10/100))
		if err != nil {
			i, err = s.Repo.OneByValue(ctx, v+(v*10/100))
		}
	}

	return i, directMatch, err
}
