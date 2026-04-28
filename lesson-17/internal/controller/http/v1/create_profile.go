package v1

import (
	"encoding/json"
	"net/http"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-17/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-17/pkg/render"
)

func (h *Handlers) CreateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := dto.CreateProfileInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Error(ctx, w, err, http.StatusBadRequest, "json decode error")

		return
	}

	output, err := h.usecase.CreateProfile(ctx, input)
	if err != nil {
		render.Error(ctx, w, err, http.StatusBadRequest, "request failed")

		return
	}

	render.JSON(w, output, http.StatusOK)
}
