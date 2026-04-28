package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	http_server "gitlab.golang-school.ru/potok-2/lessons/lesson-24/gen/http/profile_v1/server"
	ver1 "gitlab.golang-school.ru/potok-2/lessons/lesson-24/internal/controller/http/v1"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/internal/usecase"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/pkg/logger"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/pkg/metrics"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-24/pkg/otel"
)

func ProfileRouter(r *chi.Mux, uc *usecase.UseCase, m *metrics.HTTPServer) {
	v1 := ver1.New(uc)

	r.Handle("/metrics", promhttp.Handler())

	r.Route("/mnepryakhin/my-app/api", func(r chi.Router) {
		r.Use(otel.Middleware)
		r.Use(logger.Middleware)
		r.Use(metrics.NewMiddleware(m))

		r.Route("/v1", func(r chi.Router) {
			mux := http_server.NewStrictHandler(v1, []http_server.StrictMiddlewareFunc{})
			http_server.HandlerFromMux(mux, r)
		})
	})
}
