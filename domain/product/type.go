package product

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	ProductEntity struct {
		ID             primitive.ObjectID `bson:"_id,omitempty"`
		UserID         string             `bson:"user_id,omitempty"`
		Name           string             `bson:"name"`
		Category       string             `bson:"category"`
		Condition      int8               `bson:"condition"`
		Desc           string             `bson:"desc"`
		Price          float64            `bson:"price"`
		Status         int8               `bson:"status"`
		Stock          int                `bson:"stock"`
		Specifications []map[string]any   `bson:"specifications"`
	}
)
