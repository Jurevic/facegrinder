package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jurevic/facegrinder/pkg/api/v1/helper/response"
	"github.com/jurevic/facegrinder/pkg/api/v1/processor"
	"github.com/jurevic/facegrinder/pkg/datastore"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Password struct {
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Decode credentials
	var cred UserCredentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Find user
	db := datastore.DBConn()
	var user User
	err = db.Model(&user).Where("email = ?", cred.Email).Select()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Validate password
	err = checkPasswordMatch(user.Password, []byte(cred.Password))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Generate token
	jwtToken := jwt.New(jwt.SigningMethodRS256)

	// Specify claims
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = user.Id
	claims["is_superuser"] = user.IsSuperuser
	jwtToken.Claims = claims

	// Sign token
	tokenString, err := jwtToken.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	response.JsonResponse(Token{tokenString}, w)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	db := datastore.DBConn()

	// Unmarshal refresh token
	var rToken Token
	err := json.NewDecoder(r.Body).Decode(&rToken)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Parse token
	token, err := ParseToken(rToken.Token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Decode claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Find user
	user := &User{Id: claims["user_id"].(int)}
	err = db.Select(user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if !user.IsActive {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Generate token
	jwtToken := jwt.New(jwt.SigningMethodRS256)
	claims = make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = user.Id
	claims["is_superuser"] = user.IsSuperuser
	jwtToken.Claims = claims

	// Sign token
	tokenString, err := jwtToken.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	response.JsonResponse(Token{tokenString}, w)
}

func List(w http.ResponseWriter, _ *http.Request) {
	db := datastore.DBConn()

	var users []User

	err := db.Model(&users).Select()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response.JsonResponse(users, w)
}

func Retrieve(w http.ResponseWriter, r *http.Request) {
	db := datastore.DBConn()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Invalid channel id")
		return
	}

	user := &User{Id: id}
	err = db.Select(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response.JsonResponse(user, w)
}

func Create(w http.ResponseWriter, r *http.Request) {
	db := datastore.DBConn()

	// Read data
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Unmarshal password
	var password Password
	err = json.Unmarshal(data, &password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Unmarshal user data
	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	user.Password = []byte(password.Password)

	// Create user
	err = db.Insert(&user)
	if err != nil {
		panic(err)
	}

	// Create defaults
	processor.CreateDefault(user.Id)

	response.JsonResponse(user, w)
}

func Update(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//name := vars["name"]
	//
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//
	//user := new(User)
	//err = json.Unmarshal(body, user)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//
	//User.Save(name, user)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	//
	//user := new(User)
	//
	//ok := user.Remove(id)
	//if !ok {
	//	http.NotFound(w, r)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusNoContent)
}
