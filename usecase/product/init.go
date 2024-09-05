package product

import (
	"product-engine/config"
	"product-engine/domain/product"
)

type ProductUsecase struct {
	cfg           *config.Config
	productDomain product.ProductDomainItf
}

func InitProductUsecase(
	cfg *config.Config,
	productDomain product.ProductDomain,
) *ProductUsecase {
	return &ProductUsecase{
		cfg:           cfg,
		productDomain: productDomain,
	}
}
