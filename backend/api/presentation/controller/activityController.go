package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"work-management-app/application/dto/request"
	"work-management-app/application/usecase"
	"work-management-app/utility"
)

type ActivityController interface {
	AddStartWork() http.HandlerFunc
	AddEndWork() http.HandlerFunc
	AddStartBreak() http.HandlerFunc
	AddEndBreak() http.HandlerFunc
	UpdateActivity() http.HandlerFunc
	DeleteActivity() http.HandlerFunc
}

type ActivityControllerImpl struct {
	ActivityUsecase usecase.ActivityUsecase
}

func (a ActivityControllerImpl) AddStartWork() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var activity request.ActivityRequestDTO

		// bodyを取得
		if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
			log.Println("Can't get body")
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		res, err := a.ActivityUsecase.AddStarWork(&activity)
		if err != nil {
			log.Println("Error in ActivityUsecase")
			http.Error(w, `{"error":"作業開始は現在押せません"}`, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Can't encode json: %v", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}

func (a ActivityControllerImpl) AddEndWork() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var activity request.ActivityRequestDTO

		// bodyを取得
		if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
			log.Println("Can't get body")
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		res, err := a.ActivityUsecase.AddEndWork(&activity)
		if err != nil {
			log.Println("Error in ActivityUsecase")
			http.Error(w, `{"error":"作業終了は現在押せません"}`, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println("Error in json Encode")
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func (a ActivityControllerImpl) AddStartBreak() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var activity request.ActivityRequestDTO

		// bodyを取得
		if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
			log.Println("Can't get body")
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		res, err := a.ActivityUsecase.AddStartBreak(&activity)
		if err != nil {
			log.Println("Error in ActivityUsecase")
			http.Error(w, `{"error":"休憩開始は現在押せません"}`, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Can't encode json: %v", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}

func (a ActivityControllerImpl) AddEndBreak() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var activity request.ActivityRequestDTO

		// bodyを取得
		if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
			log.Println("Can't get body")
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		res, err := a.ActivityUsecase.AddEndBreak(&activity)
		if err != nil {
			log.Println("Error in ActivityUsecase")
			http.Error(w, `{"error":"休憩終了は現在押せません"}`, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println("Error in json Encode")
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func (a ActivityControllerImpl) UpdateActivity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var activity request.ActivityEditRequestDTO

		// bodyを取得
		if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// 日本時間に変換
		loc, _ := time.LoadLocation("Asia/Tokyo")
		activity.Time = activity.Time.In(loc)

		res, err := a.ActivityUsecase.Update(&activity)
		if err != nil {
			switch err.(type) {
			case utility.AuthenticationError:
				http.Error(w, err.Error(), http.StatusUnauthorized)
			case utility.NotFoundError:
				http.Error(w, err.Error(), http.StatusNotFound)
			case utility.BadRequestError:
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
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
		var activity request.ActivityDeleteRequestDTO

		// bodyを取得
		if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err := a.ActivityUsecase.DeleteByActivityID(&activity)
		if err != nil {
			switch err.(type) {
			case utility.AuthenticationError:
				http.Error(w, err.Error(), http.StatusUnauthorized)
			case utility.NotFoundError:
				http.Error(w, err.Error(), http.StatusNotFound)
			case utility.ForbiddenError:
				http.Error(w, err.Error(), http.StatusForbidden)
			case utility.BadRequestError:
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		// 成功したときのレスポンスはステータスコードだけ
		w.WriteHeader(http.StatusNoContent) // 204 No Content
	}
}

func NewActivityController(au usecase.ActivityUsecase) ActivityController {
	return &ActivityControllerImpl{ActivityUsecase: au}
}
