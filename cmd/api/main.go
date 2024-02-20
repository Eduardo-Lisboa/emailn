package main

import (
	"emailn/internal/domain/campaing"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaingService := campaing.ServiceImp{
		Repository: &database.CampaingRepository{},
	}
	handler := endpoints.Handler{
		CampaingService: &campaingService,
	}

	r.Post("/campaigns", endpoints.HandlerErro(handler.CampaingPost))
	r.Get("/campaigns", endpoints.HandlerErro(handler.CampaingGet))

	http.ListenAndServe(":3000", r)

}
