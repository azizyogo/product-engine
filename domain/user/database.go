package user

import (
	"context"
	"net/http"
	constanta "product-engine/common/const"
	"product-engine/common/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (rsc UserResource) get(ctx context.Context, filter bson.M) (UserEntity, error) {

	var res UserEntity
	err := rsc.DB.Database(constanta.DBName).Collection(constanta.CollectionUser).FindOne(ctx, filter).Decode(&res)
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
