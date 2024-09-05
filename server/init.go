package server

import (
	"context"
	"log"
	"product-engine/config"
	"product-engine/domain/product"
	"product-engine/domain/user"
	productUsecase "product-engine/usecase/product"
	userUsecase "product-engine/usecase/user"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db  *mongo.Client
	cfg *config.Config
	err error

	// domain
	productDomain product.ProductDomain
	userDomain    user.UserDomain

	// usecase
	ProductUsecase *productUsecase.ProductUsecase
	UserUsecase    *userUsecase.UserUsecase
)

func Init() error {
	cfg, err = config.LoadConfig()
	if err != nil {
		return err
	}

	var err error
	db, err = initDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return err
	}

	initLayers(db, cfg)

	return nil
}

func Close() {
	if db != nil {
		db.Disconnect(context.TODO())
	}
}

func initLayers(db *mongo.Client, cfg *config.Config) {

	// init domain
	productDomain = product.InitDomain(
		product.ProductResource{
			DB: db,
		},
	)

	userDomain = user.InitDomain(
		user.UserResource{
			DB: db,
		},
	)
	// end of init domain

	// init usecase
	ProductUsecase = productUsecase.InitProductUsecase(
		cfg,
		productDomain,
	)

	UserUsecase = userUsecase.InitUserUsecase(
		cfg,
		userDomain,
	)
	// end of init usecase
}

func initDB(cfg *config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.Mongodb.Host)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	return client, nil
}
