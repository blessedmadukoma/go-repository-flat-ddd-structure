package token

import (
	"fmt"
	"goRepositoryPattern/util"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	config *util.Config
}

func NewJWTToken(cfg *util.Config) *JWTToken {
	return &JWTToken{
		config: cfg,
	}
}

type jwtClaim struct {
	jwt.RegisteredClaims
	UserID    int64 `json:"user_id"`
	ExpiredAt int64 `json:"expired_at"`
}

// CreateToken creates a new JWT token
// func (j *JWTToken) CreateToken(user_id int64) (string, error) {
func (j *JWTToken) CreateToken(user_id int64, duration time.Duration) (string, error) {
	claims := jwtClaim{
		UserID: user_id,
		// ExpiredAt: time.Now().Add(time.Minute * 15).Unix(),
		ExpiredAt: int64(duration),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.config.SIGNINGKEY))
	if err != nil {
		return "", err
	}

	return string(tokenString), nil
}

// VerifyToken verifies a JWT token
func (j *JWTToken) VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaim{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid authentication token")
		}

		return []byte(j.config.SIGNINGKEY), nil
	})

	if err != nil {
		return 0, fmt.Errorf("invalid authentication token")
	}

	claims, ok := token.Claims.(*jwtClaim)

	if !ok {
		return 0, fmt.Errorf("invalid authentication token")
	}

	if claims.ExpiredAt < time.Now().Unix() {
		return 0, fmt.Errorf("token expired")
	}

	return claims.UserID, nil
}
