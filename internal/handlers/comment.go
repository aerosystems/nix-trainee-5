package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *BaseHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.commentRepo.FindAll()
	if err != nil {
		log.Println(err)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "...",
		Data:    comments,
	}
	_ = writeJSON(w, http.StatusOK, payload)
}

func (h *BaseHandler) GetComment(w http.ResponseWriter, r *http.Request) {
	stringID := chi.URLParam(r, "id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		_ = errorJSON(w, err, http.StatusBadRequest)
	}

	comment, err := h.commentRepo.FindByID(ID)
	if err != nil {
		log.Println(err)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "...",
		Data:    comment,
	}
	_ = writeJSON(w, http.StatusOK, payload)
}
