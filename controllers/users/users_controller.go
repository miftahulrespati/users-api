package users

import (
	"net/http"
	"strconv"

	"example.com/hello/github.com/miftahulrespati/bookstore_users-api/domain/users"
	"example.com/hello/github.com/miftahulrespati/bookstore_users-api/services"
	"example.com/hello/github.com/miftahulrespati/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) { // gin.Context only used in controllers
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return //TODO: return bad request to the caller
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)

		return
	}
	c.JSON(http.StatusCreated, user)
}
