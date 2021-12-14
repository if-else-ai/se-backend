package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	redis "github.com/go-redis/redis/v7"
	"github.com/twinj/uuid"
)

var clientRef *redis.Client

// GetRedisRef
func GetRedisRef() *redis.Client {
	if clientRef == nil {
		newClient := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDRESS"),
			Password: "",
			DB:       0,
		})

		clientRef = newClient
	}

	_, err := clientRef.Ping().Result()
	if err != nil {
		panic(err)
	}

	return clientRef
}

type TokenDetails struct {
	AccessToken string
	AccessUUID  string
	AtExpires   int64
}

// CreateToken
func CreateToken() (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = 0
	td.AccessUUID = uuid.NewV4().String()

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

// CreateAuth
func CreateAuth(userID string, td *TokenDetails) error {
	client := GetRedisRef()

	errAccess := client.Set(td.AccessUUID, userID, 0).Err()
	if errAccess != nil {
		return errAccess
	}

	return nil
}

// ExtractToken
func ExtractToken(req *http.Request) string {
	bearToken := req.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyToken
func VerifyToken(req *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(req)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

type AccessDetails struct {
	AccessUUID string
}

// ExtractTokenMetadata
func ExtractTokenMetadata(req *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(req)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		return &AccessDetails{
			AccessUUID: accessUUID,
		}, nil
	}

	return nil, err
}

// FetchAuth
func FetchAuth(authD *AccessDetails) (string, error) {
	client := GetRedisRef()

	userID, err := client.Get(authD.AccessUUID).Result()
	if err != nil {
		return "", err
	}

	return userID, nil
}

// TokenValid
func TokenValid(req *http.Request) error {
	token, err := VerifyToken(req)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

// ExtractIDFromToken
func ExtractIDFromToken(req *http.Request) (string, error) {
	tokenAuth, err := ExtractTokenMetadata(req)
	if err != nil {
		return "", err
	}

	userID, err := FetchAuth(tokenAuth)
	if err != nil {
		return "", err
	}

	return userID, nil
}

// DeleteAuth
func DeleteAuth(givenUUID string) (int64, error) {
	client := GetRedisRef()

	deleted, err := client.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}

// ExtractIDFromTokenString
func ExtractIDFromTokenString(tokenString string) (string, error) {
	tokenAuth, err := ExtractTokenMetadataFromTokenString(tokenString)
	if err != nil {
		return "", err
	}

	userID, err := FetchAuth(tokenAuth)
	if err != nil {
		return "", err
	}

	return userID, nil
}

// ExtractTokenMetadataFromTokenString
func ExtractTokenMetadataFromTokenString(tokenString string) (*AccessDetails, error) {
	token, err := VerifyTokenFromTokenString(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		return &AccessDetails{
			AccessUUID: accessUUID,
		}, nil
	}

	return nil, err
}

// VerifyTokenFromTokenString
func VerifyTokenFromTokenString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
