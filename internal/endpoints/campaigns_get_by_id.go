package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaing, err := h.CampaignService.GetBy(chi.URLParam(r, "id"))
	return campaing, 200, err

}
