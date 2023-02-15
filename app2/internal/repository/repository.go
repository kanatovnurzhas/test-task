package repository

import (
	"app2/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repository struct {
	IUserActionsRepo
}

func NewRepository(coll *mongo.Collection, ctx context.Context) *Repository {
	return &Repository{
		IUserActionsRepo: NewUserActionsRepo(coll, ctx),
	}
}

func ConnectDB(c *config.Config) (*mongo.Collection, context.Context, error) {
	ctx := context.TODO()

	mongoConn := options.Client().ApplyURI("mongodb://" + c.DbHost + ":" + c.DbPort)
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal("ошибка при подключении к mongo", err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("ошибка при проверке пинга к монго", err)
	}
	fmt.Println("Подключение к mongo успешна")

	userColl := mongoClient.Database(c.DbName).Collection("users")
	return userColl, ctx, nil
}
