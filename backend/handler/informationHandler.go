package handler

import (
	"backend/models"
	"backend/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type InformationHandler struct {
	service *services.InformationService
}

func New(s *services.InformationService) *InformationHandler {
	return &InformationHandler{
		service: s,
	}
}

func (h *InformationHandler) HandlerInformation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		information, err := h.service.GetAllInformation()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(information)

	case http.MethodPost:
		var information models.Information
		if err := json.NewDecoder(r.Body).Decode(&information); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newInformation, err := h.service.CreateNewInformation(&information)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newInformation)
	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)

	}

}

func (h *InformationHandler) HandlerInformationByParamether(w http.ResponseWriter, r *http.Request) {
	parametro := strings.TrimPrefix(r.URL.Path, "/information/")

	switch r.Method {
	case http.MethodGet:
		information, err := h.service.GetByTheme(parametro)
		if err != nil {
			http.Error(w, "No lo encontramos", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(information)
	case http.MethodPut:
		var information models.Information
		if err := json.NewDecoder(r.Body).Decode(&information); err != nil {
			http.Error(w, "input invalido", http.StatusBadRequest)
			return
		}
		parametroInt, err := strconv.Atoi(parametro)
		if err != nil {
			http.Error(w, "no se encontro", http.StatusBadRequest)
		}

		h.service.UpdateInformation(&information, parametroInt)
	case http.MethodDelete:
		parametroInt, err := strconv.Atoi(parametro)

		if err != nil {
			http.Error(w, "no se encontro", http.StatusBadRequest)
		}
		if err := h.service.DeleteInformation(parametroInt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusNoContent)

	}
}
