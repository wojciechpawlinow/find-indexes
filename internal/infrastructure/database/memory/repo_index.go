package memory

import "context"

type SliceRepository struct {
}

func (r *SliceRepository) OneByValue(ctx context.Context, v int) (int, error) {
	return 0, nil
}
