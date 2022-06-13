package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.quick.start/api"
	"net/http"
	"os"
	"strings"
)

type AuthMiddleware struct {
	Name        string
	Description string
}

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GetName returns the middleware name
func (m AuthMiddleware) GetName() string {
	return m.Name
}

// GetDescription returns the middleware description
func (m AuthMiddleware) GetDescription() string {
	return m.Description
}

func (AuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		//token := bearerToken[1]
		if len(bearerToken) < 2 {
			responseData := api.Response{
				Status:  "E001",
				Message: "Bearer token not found",
				Data:    "",
			}
			api.ErrorResponse(responseData, w)
			return
		} else {

			signedToken := bearerToken[1]

			token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			})
			//fmt.Println("jwtKey", err, token)
			if err != nil {
				responseData := api.Response{
					Status:  "E001",
					Message: "Unauthorized",
					Data:    "",
				}
				api.ErrorResponse(responseData, w)
				return
				//fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				next.ServeHTTP(w, r)
				return
			}
		}
		responseData := api.Response{
			Status:  "E001",
			Message: "Unauthorized",
			Data:    "",
		}
		api.ErrorResponse(responseData, w)
		//next.ServeHTTP(w, r)
	})
}

func Auth() AuthMiddleware {
	return AuthMiddleware{
		Name:        "Auth",
		Description: "Provides authentication over HTTP requests",
	}
}
