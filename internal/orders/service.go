package orders

import (
	"context"
	"fmt"

	repo "github.com/Sarthak-Java1124/golang-ProductionGradeAPI/internal/adapter/postgres/sqlc"
)

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}

type SVC struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &SVC{repo: repo}
}

func (s *SVC) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("Customer id cannot be empty")
	}
	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("Atleast one item is required")
	}

}
