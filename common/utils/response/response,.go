package response

import (
	"encoding/json"
	"net/http"
	"product-engine/common/model"
)

func WriteErrorResponse(w http.ResponseWriter, statusCode int, errMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := model.ErrResponse{
		Code:  statusCode,
		Cause: errMessage,
	}

	json.NewEncoder(w).Encode(response)
}

func WriteHttpResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := model.MainResponse{
		Code: statusCode,
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}
