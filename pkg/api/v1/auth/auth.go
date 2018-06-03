package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/jurevic/facegrinder/pkg/api/v1/helper/app_error"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func Init() {
	privKeyPath := viper.GetString("jwt_private_key_path")
	pubKeyPath := viper.GetString("jwt_public_key_path")

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
