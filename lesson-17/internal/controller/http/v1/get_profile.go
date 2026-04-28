package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-17/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-17/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-17/pkg/render"
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
			render.Error(ctx, w, err, http.StatusNotFound, "request failed")

		default:
			render.Error(ctx, w, err, http.StatusBadRequest, "request failed")
		}

		return
	}

	render.JSON(w, output, http.StatusOK)
}
