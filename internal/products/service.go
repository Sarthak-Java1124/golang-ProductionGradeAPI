package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) error
}

type SVC struct {
}

func NewService() Service {
	return &SVC{}
}

func (s *SVC) ListProducts(ctx context.Context) error {

}
