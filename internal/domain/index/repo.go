package index

import "context"

type Repository interface {
	OneByValue(ctx context.Context, v int) (int, int, bool, error)
}
