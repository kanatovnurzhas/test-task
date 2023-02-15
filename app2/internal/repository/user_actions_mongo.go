package repository

import (
	"app2/internal/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserActionsRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userActionsRepo struct {
	userColl *mongo.Collection
	ctx      context.Context
}

func NewUserActionsRepo(coll *mongo.Collection, ctx context.Context) IUserActionsRepo {
	return &userActionsRepo{
		userColl: coll,
		ctx:      ctx,
	}
}

func (r *userActionsRepo) CreateUser(user *models.User) error {
	_, err := r.userColl.InsertOne(r.ctx, user)
	if err != nil {
		return err
	}
	fmt.Println("user inserted successfully!")
	return nil

}

func (r *userActionsRepo) GetUserByEmail(email string) (*models.User, error) {
	filter := bson.M{"email": email}
	var user models.User
	err := r.userColl.FindOne(r.ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
