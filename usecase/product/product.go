package product

import (
	"context"
	"net/http"
	"product-engine/common/auth"
	constanta "product-engine/common/const"
	"product-engine/common/model"
	"product-engine/domain/product"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (productUC *ProductUsecase) GetProductByID(ctx context.Context, productID string) (product.ProductEntity, error) {

	objectId, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return product.ProductEntity{}, &model.ErrResponse{
			Code:  http.StatusBadRequest,
			Cause: err.Error(),
		}
	}

	resDB, err := productUC.productDomain.GetByID(ctx, objectId)
	if err != nil {
		return product.ProductEntity{}, err
	}

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)
	if claims.UserID != resDB.UserID {
		return product.ProductEntity{}, &model.ErrResponse{
			Code:  http.StatusBadRequest,
			Cause: constanta.ErrMsgUnauthorize,
		}
	}

	return resDB, nil
}

func (productUC *ProductUsecase) InsertProduct(ctx context.Context, req InsertUpdateProductReq) error {

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)

	err := productUC.productDomain.Insert(ctx, product.ProductEntity{
		UserID:         claims.UserID,
		Name:           req.Name,
		Category:       req.Category,
		Condition:      req.Condition,
		Desc:           req.Desc,
		Price:          req.Price,
		Status:         req.Status,
		Stock:          req.Stock,
		Specifications: req.Specifications,
	})
	if err != nil {
		return err
	}

	return nil
}

func (productUC *ProductUsecase) DeleteProductByID(ctx context.Context, productID string) error {

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)

	objectId, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return &model.ErrResponse{
			Code:  http.StatusBadRequest,
			Cause: err.Error(),
		}
	}

	resDB, err := productUC.productDomain.GetByID(ctx, objectId)
	if err != nil {
		return err
	}

	// Unauthorize to delete product from another user
	if resDB.UserID != claims.UserID {
		return &model.ErrResponse{
			Code:  http.StatusUnauthorized,
			Cause: constanta.ErrMsgUnauthorize,
		}
	}

	err = productUC.productDomain.Delete(ctx, objectId)
	if err != nil {
		return err
	}

	return nil
}

func (productUC *ProductUsecase) UpdateProductByID(ctx context.Context, req InsertUpdateProductReq) error {

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)

	objectId, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return &model.ErrResponse{
			Code:  http.StatusBadRequest,
			Cause: err.Error(),
		}
	}

	resDB, err := productUC.productDomain.GetByID(ctx, objectId)
	if err != nil {
		return err
	}

	// Unauthorize to delete product from another user
	if resDB.UserID != claims.UserID {
		return &model.ErrResponse{
			Code:  http.StatusUnauthorized,
			Cause: constanta.ErrMsgUnauthorize,
		}
	}

	err = productUC.productDomain.Update(ctx, product.ProductEntity{
		ID:             resDB.ID,
		Name:           req.Name,
		Category:       req.Category,
		Condition:      req.Condition,
		Desc:           req.Desc,
		Price:          req.Price,
		Status:         req.Status,
		Stock:          req.Stock,
		Specifications: req.Specifications,
	})
	if err != nil {
		return err
	}

	return nil
}

func (productUC *ProductUsecase) GetProductByName(ctx context.Context, name string) ([]product.ProductEntity, error) {

	resDB, _ := productUC.productDomain.GetByName(ctx, name)

	return resDB, nil
}
