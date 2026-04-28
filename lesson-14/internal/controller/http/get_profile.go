package http

import (
	"net/http"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/internal/dto"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-14/pkg/render"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {
	input := dto.GetProfileInput{
		ID: chi.URLParam(r, "id"),
	}

	output, err := h.profileService.GetProfile(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusOK)
}
