package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	UserDomainItf interface {
		Get(context.Context, string) (GetUserRes, error)
	}

	UserDomain struct {
		resource UserResourceItf
	}

	UserResourceItf interface {
		get(context.Context, bson.M) (UserEntity, error)
	}

	UserResource struct {
		DB *mongo.Client
	}
)

func InitDomain(rsc UserResourceItf) UserDomain {
	return UserDomain{
		resource: rsc,
	}
}

func (u UserDomain) Get(ctx context.Context, username string) (GetUserRes, error) {

	filter := bson.M{
		"username": username,
	}

	res, err := u.resource.get(ctx, filter)
	if err != nil {
		return GetUserRes{}, err
	}

	return GetUserRes{
		ID:       res.ID.Hex(),
		Username: res.Username,
		Password: res.Password,
	}, nil
}
