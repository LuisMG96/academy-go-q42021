package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/LuisMG96/academy-go-q42021/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	jwt "github.com/dgrijalva/jwt-go"
)

type Auth interface {
	GetToken(apiKey string) (string, error)
	ValidateToken(token string) (bool, error)
	ExtractTokenFromRequest(r *http.Request) (string, error)
}

//CsvService - Struct who will containt two method implementation of Csv interface
type AuthService struct {
}

func NewAuthService() *AuthService {
	authService := &AuthService{}
	return authService
}

func (a AuthService) GetToken(apiKey string) (string, error) {
	secrets, err := getSecrets()
	if err != nil {
		return "", nil
	}
	if secrets[common.GetEnvVariable("API_KEY_NAME")] != apiKey {
		return "", errors.New("401")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	signedKey := []byte(secrets["hashKey"].(string))
	tokenString, err := token.SignedString(signedKey)
	if err != nil {
		return "", errors.New("401")
	}

	return tokenString, nil

}

func (a AuthService) ValidateToken(tokenRequest string) (bool, error) {
	token, err := jwt.Parse(tokenRequest, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("401")
		}
		secrets, err := getSecrets()
		if err != nil {
			return nil, errors.New("401")
		}
		return []byte(secrets[common.GetEnvVariable("HASH_KEY_NAME")].(string)), nil
	})
	if err != nil {
		return false, errors.New("401")

	}
	if token.Valid {
		return true, nil
	} else {
		return false, errors.New("401")
	}
}

func (a AuthService) ExtractTokenFromRequest(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			token := bearerToken[1]
			return token, nil

		}
	}
	return "", errors.New("401")
}

func getSecrets() (map[string]interface{}, error) {
	secretName := common.GetEnvVariable("SECRET_NAME")
	region := common.GetEnvVariable("REGION")

	//Create a Secrets Manager client
	sess, err := session.NewSession()
	if err != nil {
		// Handle session creation error
		return nil, errors.New("401")
	}
	svc := secretsmanager.New(sess,
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {

		return nil, errors.New("401")
	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	}
	secrets := make(map[string]interface{})
	err = json.Unmarshal([]byte(secretString), &secrets)
	if err != nil {
		return nil, errors.New("401")
	}

	return secrets, nil
	// Your code goes here.
}
