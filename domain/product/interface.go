package product

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ProductDomainItf interface {
		GetByID(context.Context, primitive.ObjectID) (ProductEntity, error)
		GetByName(context.Context, string) ([]ProductEntity, error)
		Insert(context.Context, ProductEntity) error
		Delete(context.Context, primitive.ObjectID) error
		Update(context.Context, ProductEntity) error
	}

	ProductDomain struct {
		resource ProductResourceItf
	}

	ProductResourceItf interface {
		getByIDDB(context.Context, bson.M) (ProductEntity, error)
		getByNameDB(context.Context, bson.M) ([]ProductEntity, error)
		insertDB(context.Context, ProductEntity) error
		deleteDB(context.Context, bson.M) error
		updateDB(context.Context, bson.M, bson.M) error
	}

	ProductResource struct {
		DB *mongo.Client
	}
)

func InitDomain(rsc ProductResourceItf) ProductDomain {
	return ProductDomain{
		resource: rsc,
	}
}

func (p ProductDomain) GetByID(ctx context.Context, id primitive.ObjectID) (ProductEntity, error) {

	filter := bson.M{
		"_id": id,
	}

	return p.resource.getByIDDB(ctx, filter)
}

func (p ProductDomain) GetByName(ctx context.Context, name string) ([]ProductEntity, error) {

	filter := bson.M{
		"name": bson.M{
			"$regex":   name,
			"$options": "i", // case-insensitive search
		},
	}

	return p.resource.getByNameDB(ctx, filter)
}

func (p ProductDomain) Insert(ctx context.Context, req ProductEntity) error {

	return p.resource.insertDB(ctx, req)
}

func (p ProductDomain) Delete(ctx context.Context, id primitive.ObjectID) error {

	selector := bson.M{
		"_id": id,
	}

	return p.resource.deleteDB(ctx, selector)
}

func (p ProductDomain) Update(ctx context.Context, req ProductEntity) error {

	selector := bson.M{
		"_id": req.ID,
	}

	values := bson.M{
		"$set": req,
	}

	return p.resource.updateDB(ctx, selector, values)
}
