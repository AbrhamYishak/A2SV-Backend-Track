package Repositories

import (
	"context"
	"errors"
	"task7/Domain"
	"task7/Usecases"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
type UserRepo struct {
	Collection *mongo.Collection
	Context    context.Context

}
func NewUserRepo(u *mongo.Collection) Usecases.UserRepoI {
	col := u 
	ctx := context.Background()

	return &UserRepo{
		Collection: col,
		Context:    ctx,
	}
}
func (ur *UserRepo) UserExist (username string)(Domain.User, bool){ 
	var u Domain.User
    if err := ur.Collection.FindOne(ur.Context, bson.M{"username": u.Username}).Decode(&u); err == nil{
		return u,false
	}
	return u,true
}
func (ur *UserRepo) CountUsers () (int64,error){
	count, err := ur.Collection.CountDocuments(ur.Context,bson.M{})
	if err != nil{
        return 0,errors.New("could not find the number of users")
	}
	return count, nil
}
func (ur *UserRepo) Register(u * Domain.User) error{
    u.ID = primitive.NewObjectID()
	_, err := ur.Collection.InsertOne(ur.Context, u)
	if err != nil {
		return errors.New("failed to create user")
	}
    return nil
}
func (ur *UserRepo) Isadmin(username string)(bool, error){
    user, exists := ur.UserExist(username) 
	if !exists{
		return false,errors.New("user does not exist")
	}
	return user.Isadmin, nil
}
