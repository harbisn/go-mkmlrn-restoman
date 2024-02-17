package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type Error struct {
	StatusCode int    `json:"errorCode"`
	Message    string `json:"errorMessage"`
}

func WriteFailResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonData, _ := json.Marshal(Error{
		StatusCode: statusCode,
		Message:    message,
	})
	_, err := w.Write(jsonData)
	if err != nil {
		return
	}
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}, headMap map[string]string) {
	w.Header().Add("Content-Type", "application/json")
	if headMap != nil && len(headMap) > 0 {
		for key, val := range headMap {
			w.Header().Add(key, val)
		}
	}
	w.WriteHeader(statusCode)
	jsonData, _ := json.Marshal(data)
	_, err := w.Write(jsonData)
	if err != nil {
		return
	}
}

func GetIdFromPath(r *http.Request) (uint64, error) {
	vars := mux.Vars(r)
	strId := vars["id"]
	intId, err := strconv.Atoi(strId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to Get ID from path variable")
		return 0, err
	}
	return uint64(intId), nil
}
