package logger

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/router"
)

func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := router.WriterWrapper(w)
		next.ServeHTTP(ww, r.WithContext(r.Context()))

		log.Info().
			Str("proto", "http").
			Int("code", ww.Code()).
			Str("method", fmt.Sprintf("%s %s", r.Method, router.ExtractPath(r.Context()))).
			Send()
	}

	return http.HandlerFunc(fn)
}
