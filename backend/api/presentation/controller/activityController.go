package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/yoshimi-I/AttendanceApp/usecase"
	"github.com/yoshimi-I/AttendanceApp/usecase/dto/request"
	"net/http"
	"strconv"
)

type ActivityController interface {
	AddStartActivity() http.HandlerFunc
	AddEndActivity() http.HandlerFunc
	UpdateActivity() http.HandlerFunc
	DeleteActivity() http.HandlerFunc
}

type ActivityControllerImpl struct {
	ActivityUsecase usecase.ActivityUsecase
}

func (a ActivityControllerImpl) AddStartActivity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var activity request.ActivityRequestDTO

		// bodyを取得
		if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		res, err := a.ActivityUsecase.AddStartActivity(&activity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}

func (a ActivityControllerImpl) AddEndActivity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var activity request.ActivityRequestDTO

		// bodyを取得
		if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// IDをURLから取得
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		res, err := a.ActivityUsecase.AddEndActivity(&activity, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func (a ActivityControllerImpl) UpdateActivity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var activity request.ActivityRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// IDをURLから取得
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		res, err := a.ActivityUsecase.Update(&activity, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func (a ActivityControllerImpl) DeleteActivity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// IDをURLから取得
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		err = a.ActivityUsecase.DeleteByActivityID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent) // 204 No Content
	}
}

func NewActivityController(au usecase.ActivityUsecase) ActivityController {
	return &ActivityControllerImpl{ActivityUsecase: au}
}
