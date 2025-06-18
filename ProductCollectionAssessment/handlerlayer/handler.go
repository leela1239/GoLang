package handlerlayer

import (
	"encoding/json"
	"log"
	"net/http"
	"product-details/models"
	"product-details/servicelayer"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *servicelayer.Service
}

func NewHandler(s *servicelayer.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	getquery := mux.Vars(r)
	product, err := h.service.GetProducts(getquery)
	log.Print(product, err)
}

func (h *Handler) PostProducts(w http.ResponseWriter, r *http.Request) {
	var products models.Product
	json.NewDecoder(r.Body).Decode(&products)
	h.service.PostProducts(&products)
}

func (h *Handler) PutProducts(w http.ResponseWriter, r *http.Request) {

	pid := r.PathValue("id")
	var products models.Product
	h.service.PutProducts(pid, &products)
}

func (h *Handler) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	pid := r.PathValue("id")
	vid := r.URL.Query().Get("id")
	h.service.DeleteProducts(pid, vid)
}
