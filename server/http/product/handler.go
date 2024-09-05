package product

import (
	"encoding/json"
	"net/http"
	"product-engine/common/model"
	"product-engine/common/utils/response"
	"product-engine/server"
	"product-engine/usecase/product"

	"github.com/gorilla/mux"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.MainResponse{
		Code: 200,
		Data: struct {
			Msg string
		}{
			Msg: "PONG",
		},
	})
}

func HandleGetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqBody := product.GetByIDReq{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := server.ProductUsecase.GetProductByID(ctx, reqBody.ProductID)
	if customErr, ok := err.(*model.ErrResponse); ok {
		response.WriteErrorResponse(w, customErr.Code, customErr.Cause)
		return
	}

	response.WriteHttpResponse(w, http.StatusOK, res)
}

func HandleInsertProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqBody := product.InsertUpdateProductReq{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := server.ProductUsecase.InsertProduct(ctx, reqBody)
	if customErr, ok := err.(*model.ErrResponse); ok {
		response.WriteErrorResponse(w, customErr.Code, customErr.Cause)
		return
	}

	response.WriteHttpResponse(w, http.StatusCreated, nil)
}

func HandleDeleteByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	err := server.ProductUsecase.DeleteProductByID(ctx, vars["productID"])
	if customErr, ok := err.(*model.ErrResponse); ok {
		response.WriteErrorResponse(w, customErr.Code, customErr.Cause)
		return
	}

	response.WriteHttpResponse(w, http.StatusOK, nil)
}

func HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqBody := product.InsertUpdateProductReq{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := server.ProductUsecase.UpdateProductByID(ctx, reqBody)
	if customErr, ok := err.(*model.ErrResponse); ok {
		response.WriteErrorResponse(w, customErr.Code, customErr.Cause)
		return
	}

	response.WriteHttpResponse(w, http.StatusOK, nil)
}
