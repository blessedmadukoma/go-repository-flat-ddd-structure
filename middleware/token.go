package middleware

import (
	"errors"
	"fmt"
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/messages"
	"goRepositoryPattern/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// authMiddleware authorizes/validates a user
func (m Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var R = messages.ResponseFormat{}
		var err error

		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err = errors.New("authorization header not provided")

			R.Error = append(R.Error, err.Error())
			R.Message = messages.UnAuthorizedAccess
			ctx.JSON(messages.Response(http.StatusUnauthorized, R))
			ctx.Abort()
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err = errors.New("invalid authorization header format")

			R.Error = append(R.Error, err.Error())
			R.Message = messages.UnAuthorizedAccess
			ctx.JSON(messages.Response(http.StatusUnauthorized, R))
			ctx.Abort()
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err = fmt.Errorf("unsupported authorization type %s", authorizationType)
			R.Error = append(R.Error, err.Error())
			R.Message = messages.UnAuthorizedAccess
			ctx.JSON(messages.Response(http.StatusUnauthorized, R))
			ctx.Abort()
			return
		}

		accessToken := fields[1]
		payload, err := m.repo.Token.VerifyToken(accessToken)
		if err != nil {
			R.Error = append(R.Error, err.Error())
			R.Message = messages.ErrInvalidToken.Error()
			ctx.JSON(messages.Response(http.StatusUnauthorized, R))
			ctx.Abort()
			return
		}

		// store payload in the context
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

func (m Middleware) AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var R = messages.ResponseFormat{}
		var err error

		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err = errors.New("authorization header not provided")

			R.Error = append(R.Error, err.Error())
			R.Message = messages.UnAuthorizedAccess
			ctx.JSON(messages.Response(http.StatusUnauthorized, R))
			ctx.Abort()
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err = errors.New("invalid authorization header format")

			R.Error = append(R.Error, err.Error())
			R.Message = messages.UnAuthorizedAccess
			ctx.JSON(messages.Response(http.StatusUnauthorized, R))
			ctx.Abort()
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err = fmt.Errorf("unsupported authorization type %s", authorizationType)
			R.Error = append(R.Error, err.Error())
			R.Message = messages.UnAuthorizedAccess
			ctx.JSON(messages.Response(http.StatusUnauthorized, R))
			ctx.Abort()
			return
		}

		accessToken := fields[1]
		payload, err := m.repo.Token.VerifyToken(accessToken)
		if err != nil {
			R.Error = append(R.Error, err.Error())
			R.Message = messages.ErrInvalidToken.Error()
			ctx.JSON(messages.Response(http.StatusUnauthorized, R))
			ctx.Abort()
			return
		}

		a, err := m.repo.DB.GetAccountByID(ctx, payload)
		if err != nil {
			R.Error = append(R.Error, err.Error())
			R.Message = messages.SomethingWentWrong
			ctx.JSON(messages.Response(http.StatusInternalServerError, R))
			ctx.Abort()
			return
		}

		if a.Role != database.RoleAdmin {
			R.Error = append(R.Error, messages.UnAuthorizedAccess)
			R.Message = messages.UnAuthorizedAccess
			ctx.JSON(messages.Response(http.StatusUnauthorized, R))
			ctx.Abort()
			return
		}

		// store payload in the context
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

// getAuthorizationPayload retrieves the authorization payload from the context
func getAuthorizationPayload(ctx *gin.Context) (*token.JWTToken, error) {
	payload, ok := ctx.Get(authorizationPayloadKey)
	if !ok {
		return nil, errors.New("authorization payload not found")
	}

	return payload.(*token.JWTToken), nil
}
