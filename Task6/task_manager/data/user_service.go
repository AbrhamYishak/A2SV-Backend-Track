package data

import (
	"context"
	"errors"
	"net/http"
	"task6/models"
	"time"
    "github.com/golang-jwt/jwt/v5"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"	
)
var userCollection *mongo.Collection 		
func InitUserController( users *mongo.Collection){
	userCollection = users
}
type Claims struct {
	Username string
	Userid string
	Isadmin bool
	jwt.RegisteredClaims
}
func GetToken(username string,userid string, admin bool)(string, error){
	var jwtKey = []byte("your_jwt_secret")
	expirationTime := time.Now().AddDate(0, 0, 1)
	claims := &Claims{
		Username : username,
		Isadmin : admin,
		Userid : userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func VerifyToken(requestToken string, secret string)(bool,error){
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
func ExtractFromToken(tokenString string, secret string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username, nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}
func Register(u models.User) (int, error){
    u.ID = primitive.NewObjectID()
	var test models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
    if err := userCollection.FindOne(ctx, bson.M{"username": u.Username}).Decode(&test); err == nil{
		return http.StatusBadRequest, errors.New("username already in use find another username")
	}
	count, err := userCollection.CountDocuments(ctx,bson.M{})
	if err != nil{
        return http.StatusInternalServerError, errors.New("could not find the number of users")
	}
	if count == 0{
		u.Isadmin = true
	}
	_, err = userCollection.InsertOne(ctx, u)
	if err != nil {
		return http.StatusInternalServerError, errors.New("failed to create user")
	}
    return http.StatusCreated, nil
}
func Login(u models.User) (int,string, error){
    var user models.User 
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := userCollection.FindOne(ctx, bson.M{"username": u.Username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
		 	return http.StatusNotFound,"", errors.New("could not find user with this username")
		}
		return http.StatusInternalServerError,"", errors.New("error occur in the database")
	}
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil{
	   return http.StatusBadRequest,"", errors.New("wrong password")
	}
	token, err := GetToken(user.Username,user.ID.Hex(), user.Isadmin)
	if err != nil{
		return http.StatusInternalServerError,"", errors.New("could not generate token")
	}
	return http.StatusOK,token,nil
    
}
func Promote(u models.User) (int, error){
    var user models.User 
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := userCollection.FindOne(ctx, bson.M{"username": u.Username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
		 	return http.StatusNotFound, errors.New("could not find user with this username")
		}
		return http.StatusInternalServerError, errors.New("error occur in the database")
	}
    update := bson.M{
		"$set": bson.M{
			"username":  user.Username,
			"password":  user.Password,
			"isadmin":   true,
		},
	}
	result, err := userCollection.UpdateOne(ctx, bson.M{"username": user.Username}, update)
	if err != nil || result.MatchedCount == 0 {
		return http.StatusInternalServerError, errors.New("could promote the user")
	}
	return http.StatusOK, nil
}
func Isadmin(username string)(bool,error){
    var user models.User 
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
		 	return false, errors.New("could not find user with this username")
		}
		return false, errors.New("error occur in the database")
	}
	return user.Isadmin, nil
}
