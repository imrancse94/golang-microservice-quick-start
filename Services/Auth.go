package Services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.quick.start/models"
	"io"
	"os"
	"strconv"
	"time"
)

type AuthUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Jwt struct {
}

func (j Jwt) CreateToken(user models.User) (models.Token, error) {
	var err error
	expire, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRES_IN"), 10, 64)
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(time.Duration(expire))).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken := models.Token{}

	jwtToken.AccessToken, err = token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return jwtToken, err
	}
	id := int(user.ID)
	permissionData := models.GetRolePageByUserId(id)
	jwtToken.User = user
	jwtToken.Permissions = permissionData.Data
	jwtToken.Expire = expire * 60
	return j.createRefreshToken(jwtToken)
}

func (Jwt) ValidateToken(accessToken string) (AuthUser, error) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	user := AuthUser{}
	//v, _ := err.(*jwt.ValidationError)
	fmt.Println("sssss", jwt.ValidationErrorExpired)
	if err != nil {
		return user, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		user.ID = int(payload["user_id"].(float64))
		user.Email = payload["email"].(string)
		return user, nil
	}

	return user, errors.New("invalid token")
}

func (Jwt) ValidateRefreshToken(model models.Token) (models.User, error) {
	sha1 := sha1.New()
	io.WriteString(sha1, os.Getenv("JWT_SECRET_KEY"))

	user := models.User{}
	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return user, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return user, err
	}

	data, err := base64.URLEncoding.DecodeString(model.RefreshToken)
	if err != nil {
		return user, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return user, err
	}

	if string(plain) != model.AccessToken {
		return user, errors.New("invalid token")
	}

	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(model.AccessToken, claims)

	if err != nil {
		return user, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return user, errors.New("invalid token")
	}

	user.ID = payload["user_id"].(uint)

	return user, nil
}

func (Jwt) createRefreshToken(token models.Token) (models.Token, error) {
	sha1 := sha1.New()
	io.WriteString(sha1, os.Getenv("JWT_SECRET_KEY"))

	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		fmt.Println(err.Error())

		return token, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return token, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return token, err
	}

	token.RefreshToken = base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(token.AccessToken), nil))

	return token, nil
}
func GetAuth() {

}
