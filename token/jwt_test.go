package token

import (
	"testing"
	"time"

	"goRepositoryPattern/messages"
	"goRepositoryPattern/util"

	// "go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

// var tokenController *JWTToken
// var tokenController *NewJWTToken(&utils.Config{})
var tokenController = NewJWTToken(&util.Config{})

func TestJWTMaker(t *testing.T) {
	// userId := primitive.NewObjectID()
	userId := int64(1)

	// token, err := tokenController.CreateToken(userId, time.Minute)
	token, err := tokenController.CreateToken(userId, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	tokenValue, err := tokenController.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, tokenValue)

	require.NotZero(t, tokenValue)
}

func TestExpiredJWTToken(t *testing.T) {
	// maker, err := NewJWTMaker(util.RandomString(32))
	// require.NoError(t, err)

	// userId := primitive.NewObjectID()
	userId := int64(1)

	tokenString, err := tokenController.CreateToken(userId, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	tokenValue, err := tokenController.VerifyToken(tokenString)
	require.Error(t, err)
	// require.EqualError(t, err, ErrExpiredToken.Error())
	require.EqualError(t, err, messages.ErrExpiredToken.Error())
	require.Zero(t, tokenValue)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	// userId := primitive.NewObjectID()
	userId := int64(1)

	claims := jwtClaim{
		UserID:    userId,
		ExpiredAt: time.Now().Add(time.Minute * 15).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	tokenString, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType) // UnsafeAllowNoneSignatureType -> used only for testing, not in production
	require.NoError(t, err)

	value, err := tokenController.VerifyToken(tokenString)
	require.Error(t, err)
	// require.EqualError(t, err, ErrInvalidToken.Error())
	require.EqualError(t, err, messages.ErrInvalidToken.Error())
	require.Zero(t, value)
}
