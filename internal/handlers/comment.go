package handlers

import (
	"log"
	"net/http"

	"github.com/aerosystems/nix-trainee-4/internal/models"
)

func (h *BaseHandler) GetComment(w http.ResponseWriter, r *http.Request) {
	type responseData struct {
		status string
		result models.Comment
	}

	comment, err := h.commentRepo.FindByID(301)
	if err != nil {
		log.Println(err)
	}

	result := responseData{
		status: "OK",
		result: *comment,
	}

	payload := jsonResponse{
		Error:   false,
		Message: "...",
		Data:    result,
	}
	_ = writeJSON(w, http.StatusOK, payload)
}
