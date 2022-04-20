package handler

import (
	"kube2/internal/models"
	"net/http"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := &models.Response{Status: "OK"}

		bt, err := response.Marshal()
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(bt)
	}
}
