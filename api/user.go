package api

import (
	"goRepositoryPattern/messages"
	"log"

	"github.com/gin-gonic/gin"
)

func (c Controller) GetUserByID(ctx *gin.Context) {
	R := messages.ResponseFormat{}

	id := "2"
	email, err := c.services.UserService.GetUserByID(ctx, id)
	if err != nil {
		log.Fatal("error:", err)

	}

	log.Println("email:", email)

	// return email, nil

	R.Data = gin.H{
		"email": email,
	}

	ctx.JSON(messages.Response(0, R))
}
