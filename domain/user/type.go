package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	UserEntity struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Username string             `bson:"username,omitempty"`
		Password string             `bson:"password,omitempty"`
	}

	GetUserRes struct {
		ID       string
		Username string
		Password string
	}
)
