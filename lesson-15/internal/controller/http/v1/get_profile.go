package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-15/pkg/render"
)

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := dto.GetProfileInput{
		ID: chi.URLParam(r, "id"),
	}

	output, err := h.usecase.GetProfile(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			render.Error(w, err, http.StatusNotFound, "request failed")

		default:
			render.Error(w, err, http.StatusBadRequest, "request failed")
		}

		return
	}

	render.JSON(w, output, http.StatusOK)
}
