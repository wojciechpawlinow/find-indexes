package service

import (
	"context"
	"errors"

	"github.com/wojciechpawlinow/find-indexes/internal/domain/index"
)

type IndexPort interface {
	FindByValue(ctx context.Context, v int) (int, int, bool, error)
}

type IndexService struct {
	Repo index.Repository
}

var _ IndexPort = (*IndexService)(nil)

// FindByValue matches the exact value in the storage or the closest one
// v - value to find
// returns:
//
//	idx - index of the value,
//	directMatch - true if the value was found, false if the closest value was returned,
//	err - error message.
func (s *IndexService) FindByValue(ctx context.Context, v int) (int, int, bool, error) {

	if v < 0 {
		return -1, -1, false, errors.New("invalid value range")
	}

	return s.Repo.OneByValue(ctx, v)
}
