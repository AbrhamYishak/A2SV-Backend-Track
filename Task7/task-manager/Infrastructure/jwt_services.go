package Infrastructure
import (
	"task7/Usecases"
	"task7/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
    "github.com/golang-jwt/jwt/v5"
	"fmt"
)
type JwtService struct{
}
func NewJwtService() Usecases.JwtI{
    return &JwtService{}
}
type Claims struct {
	Username string
	Userid primitive.ObjectID 
	jwt.RegisteredClaims
}
func (js *JwtService) GenerateToken(user *Domain.User)(string, error){
	var jwtKey = []byte(Env.Jwt_Secret)
	expirationTime := time.Now().AddDate(0, 0, 1)
	claims := &Claims{
		Username : user.Username,
		Userid : user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func(js *JwtService) VerifyToken(requestToken string)(bool,error){
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Env.Jwt_Secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
func(js *JwtService) ExtractFromToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Env.Jwt_Secret), nil
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
