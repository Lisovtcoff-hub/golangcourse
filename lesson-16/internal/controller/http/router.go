package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	ver1 "gitlab.golang-school.ru/potok-2/lessons/lesson-16/internal/controller/http/v1"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/internal/usecase"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/logger"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/metrics"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/otel"
)

func ProfileRouter(r *chi.Mux, uc *usecase.UseCase, m *metrics.HTTPServer) {
	v1 := ver1.New(uc)

	r.Handle("/metrics", promhttp.Handler())

	r.Route("/mnepryakhin/my-app/api", func(r chi.Router) {
		r.Use(logger.Middleware)
		r.Use(metrics.NewMiddleware(m))
		r.Use(otel.Middleware)

		r.Route("/v1", func(r chi.Router) {
			r.Post("/profile", v1.CreateProfile)
			r.Put("/profile", v1.UpdateProfile)
			r.Get("/profile/{id}", v1.GetProfile)
			r.Delete("/profile/{id}", v1.DeleteProfile)
		})
	})
}
