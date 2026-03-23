package products

import (
	"context"
	"log"

	repo "github.com/Sarthak-Java1124/golang-ProductionGradeAPI/internal/adapter/postgres/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
}

type SVC struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &SVC{repo: repo}
}

func (s *SVC) ListProducts(ctx context.Context) ([]repo.Product, error) {
	products, err := s.repo.ListProducts(ctx)
	if err != nil {
		log.Println("The error in list products is : ", err)
		return products, err

	}
	return products, nil

}
