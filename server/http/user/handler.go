package user

import (
	"encoding/json"
	"net/http"
	"product-engine/common/utils/response"
	"product-engine/server"
	"product-engine/usecase/user"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqBody := user.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := server.UserUsecase.Login(ctx, reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteHttpResponse(w, http.StatusOK, res)
}
