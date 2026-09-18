package handler

import (
	"backend/models"
	"backend/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ScheduleHandler struct {
	services *services.ScheduleService
}

func NewScheduleHandler(s *services.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		services: s}
}

func (h *ScheduleHandler) ScheduleInformation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//peticion tipo get
	case http.MethodGet:
		scheduleInformation, err := h.services.GetAllInformationSchedule()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(scheduleInformation)

	case http.MethodPost:
		var informationScheduler models.Schedule
		if err := json.NewDecoder(r.Body).Decode(&informationScheduler); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newInformationScheduler, err := h.services.CreateInformationSchedule(&informationScheduler)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newInformationScheduler)
	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)

	}
}

func (h *ScheduleHandler) HandlerInformationByParamether(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/schedule/")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Input invalido", http.StatusBadRequest)
	}

	switch r.Method {
	case http.MethodPut:
		var scheduleInformation models.Schedule

		if err := json.NewDecoder(r.Body).Decode(&scheduleInformation); err != nil {
			http.Error(w, "input invalido", http.StatusBadRequest)
			return
		}
		h.services.UpdateInformationSchedule(&scheduleInformation, idInt)

	case http.MethodDelete:
		if err := h.services.DeleteInformationSchedule(idInt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusNoContent)
	}

}
