package products

import (
	"log"
	"net/http"

	JSON "github.com/Sarthak-Java1124/golang-ProductionGradeAPI/internal/json"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	JSON.Write(w, http.StatusOK, products)
}
func (h *Handler) FindProductsById(w http.ResponseWriter, r *http.Request) {

}
