package memory

import (
	"context"

	"github.com/wojciechpawlinow/find-indexes/internal/domain/index"
)

type SliceRepository struct {
}

var _ index.Repository = (*SliceRepository)(nil)

func (r *SliceRepository) OneByValue(ctx context.Context, v int) (int, error) {
	return 0, nil
}
