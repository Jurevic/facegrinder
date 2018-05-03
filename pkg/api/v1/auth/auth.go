package auth

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/jurevic/facegrinder/pkg/api/v1/helper/app_error"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	privateKeyPath = "JWT_PRIVATE_KEY_PATH"
	publicKeyPath  = "JWT_PUBLIC_KEY_PATH"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func Init() {
	privKeyPath, ok := os.LookupEnv(privateKeyPath)
	if !ok {
		panic(fmt.Sprintf("%s environment variable required but not set", privateKeyPath))
	}

	pubKeyPath, ok := os.LookupEnv(publicKeyPath)
	if !ok {
		panic(fmt.Sprintf("%s environment variable required but not set", publicKeyPath))
	}

	signBytes, err := ioutil.ReadFile(privKeyPath)
	app_error.Fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	app_error.Fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	app_error.Fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	app_error.Fatal(err)
}

func ValidateToken(r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	return token, err
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	return token, err
}
