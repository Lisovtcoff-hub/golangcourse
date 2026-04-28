package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/render"
)

func (h *Handlers) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := dto.DeleteProfileInput{
		ID: chi.URLParam(r, "id"),
	}

	err := h.usecase.DeleteProfile(ctx, input)
	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
