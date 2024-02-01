package main

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaing"
	"emailn/internal/infrastructure/database"
	internalerrors "emailn/internal/internalErrors"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaing.Service{
		Repository: &database.CampaingRepository{},
	}

	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		request := contract.NewCampaing{}
		render.DecodeJSON(r.Body, &request)

		id, err := service.Create(request)
		if err != nil {

			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, 500)
				render.JSON(w, r, map[string]string{"error": err.Error()})

			} else {

				render.Status(r, 400)
				render.JSON(w, r, map[string]string{"error": err.Error()})
			}

			return
		}
		render.Status(r, 201)
		render.JSON(w, r, map[string]string{"id": id})
	})

	http.ListenAndServe(":3000", r)

}
