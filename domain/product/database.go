package product

import (
	"context"
	"net/http"
	constanta "product-engine/common/const"
	"product-engine/common/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (rsc ProductResource) getByIDDB(ctx context.Context, filter bson.M) (ProductEntity, error) {

	var res ProductEntity
	err := rsc.DB.Database(constanta.DBName).Collection(constanta.CollectionProduct).FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return res, &model.ErrResponse{
				Code:  http.StatusNotFound,
				Cause: constanta.ErrMsgNotFound,
			}
		}
		return res, &model.ErrResponse{
			Code:  http.StatusInternalServerError,
			Cause: constanta.ErrMsgInternalError,
		}
	}

	return res, nil
}

func (rsc ProductResource) getByNameDB(ctx context.Context, filter bson.M) ([]ProductEntity, error) {

	var res []ProductEntity

	cursor, err := rsc.DB.Database(constanta.DBName).Collection(constanta.CollectionProduct).Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return res, &model.ErrResponse{
				Code:  http.StatusNotFound,
				Cause: constanta.ErrMsgNotFound,
			}
		}
		return res, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var record ProductEntity
		if err := cursor.Decode(&record); err != nil {
			return nil, err
		}
		res = append(res, record)
	}

	if err := cursor.Err(); err != nil {
		return res, &model.ErrResponse{
			Code:  http.StatusInternalServerError,
			Cause: constanta.ErrMsgInternalError,
		}
	}

	return res, nil
}

func (rsc ProductResource) insertDB(ctx context.Context, doc ProductEntity) error {

	_, err := rsc.DB.Database(constanta.DBName).Collection(constanta.CollectionProduct).InsertOne(ctx, doc)
	if err != nil {
		return &model.ErrResponse{
			Code:  http.StatusInternalServerError,
			Cause: constanta.ErrMsgInternalError,
		}
	}

	return nil
}

func (rsc ProductResource) deleteDB(ctx context.Context, doc bson.M) error {

	_, err := rsc.DB.Database(constanta.DBName).Collection(constanta.CollectionProduct).DeleteOne(ctx, doc)
	if err != nil {
		return &model.ErrResponse{
			Code:  http.StatusInternalServerError,
			Cause: constanta.ErrMsgInternalError,
		}
	}

	return nil
}

func (rsc ProductResource) updateDB(ctx context.Context, selector, values bson.M) error {

	_, err := rsc.DB.Database(constanta.DBName).Collection(constanta.CollectionProduct).UpdateOne(ctx, selector, values)
	if err != nil {
		return &model.ErrResponse{
			Code:  http.StatusInternalServerError,
			Cause: constanta.ErrMsgInternalError,
		}
	}

	return nil
}
