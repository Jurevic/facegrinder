package middleware

import (
	"fmt"
	"github.com/jurevic/facegrinder/pkg/api/v1/auth"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"context"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.ValidateToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access to this resource")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !token.Valid || !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
			return
		}

		r = addUserContext(r, claims)

		next.ServeHTTP(w, r)
	})
}

func addUserContext(r *http.Request, claims jwt.MapClaims) *http.Request {
	ctx := context.WithValue(r.Context(), "user_id", int(claims["user_id"].(float64)))
	ctx = context.WithValue(ctx, "is_superuser", claims["is_superuser"])

	return r.WithContext(ctx)
}
