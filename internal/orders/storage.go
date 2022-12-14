package orders

import "context"

type Repository interface {
	GetServiceTotal(ctx context.Context, year, month string) ([]ServiceTotal, error)
}
