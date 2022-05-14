package Services

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	"go.quick.start/app/requests"
	"go.quick.start/models"
	"os"
	"strconv"
	"time"
)

// Create a dummy local db instance as a key value pair
var userdb = map[string]string{
	"imrancse94@gmail.com": "Nop@ss1234",
}

// Login user login function
func Login(input requests.Credential) (string string, data interface{}, error string) {
	expire, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRES_IN"), 10, 64)
	user := models.GetUserByEmail(input.Email)

	fmt.Println("user111", user.Data)
	userPassword, ok := userdb[input.Email]

	// if user exist, verify the password
	if !ok || userPassword != input.Password {
		//w.WriteHeader(http.StatusUnauthorized)
		return "", user.Data, "Invalid email or password"
	}
	// Create a token object and add the Username and StandardClaims
	expires := time.Now().Add(time.Duration(time.Duration(expire) * time.Minute)).Unix()
	var tokenClaim = requests.Token{
		Email: input.Email,
		StandardClaims: jwt.StandardClaims{
			// Enter expiration in milisecond
			ExpiresAt: expires,
		},
	}

	// Create a new claim with HS256 algorithm and token claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, user.Data, ""
}
