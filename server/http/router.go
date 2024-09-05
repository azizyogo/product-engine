package http

import (
	constanta "product-engine/common/const"
	"product-engine/middleware"
	"product-engine/server/http/product"
	"product-engine/server/http/user"

	"github.com/gorilla/mux"
)

func NewServer() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", product.HandlePing).Methods(constanta.HTTPGet)

	// Auth
	// This endpoinst doesn't supposed to be here, this endpoint for support testing only
	router.HandleFunc("/login", user.HandleLogin).Methods("POST")

	// Product Management
	// Get product details
	router.HandleFunc("/product", middleware.Authorize(product.HandleGetByID)).Methods(constanta.HTTPGet)
	// Insert new product
	router.HandleFunc("/product", middleware.Authorize(product.HandleInsertProduct)).Methods(constanta.HTTPPost)
	// Delete existed product
	router.HandleFunc("/product/{productID}", middleware.Authorize(product.HandleDeleteByID)).Methods(constanta.HTTPDelete)
	// Update product details
	router.HandleFunc("/product", middleware.Authorize(product.HandleUpdateByID)).Methods(constanta.HTTPPut)

	return router
}
