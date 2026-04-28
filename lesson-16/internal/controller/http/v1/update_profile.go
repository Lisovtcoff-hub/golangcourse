package v1

import (
	"encoding/json"
	"net/http"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/render"
)

// UpdateProfile обновляет существующий или создаёт новый профиль.
func (h *Handlers) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := dto.UpdateProfileInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "json decode error")

		return
	}

	err = h.usecase.UpdateProfile(ctx, input)
	if err != nil {
		render.Error(w, err, http.StatusBadRequest, "request failed")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
