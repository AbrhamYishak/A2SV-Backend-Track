package Repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"task7/Domain"
	"task7/Usecases"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type UserRepo struct {
	Collection *mongo.Collection
	Context    context.Context
}
func NewUserRepo() Usecases.UserRepoI {
	col := initDB()
	ctx := context.Background()

	return &UserRepo{
		Collection: col,
		Context:    ctx,
	}
}
func (us *UserRepo) Register(u * Domain.User) error{
     
}
