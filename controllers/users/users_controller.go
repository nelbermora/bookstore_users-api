package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nelbermora/bookstore_users-api/domain/users"
	"github.com/nelbermora/bookstore_users-api/services"
	"github.com/nelbermora/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var newUser users.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		restErr := errors.NewBadRequestErr("Invalid Json")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	result, saveErr := services.CreateUser(newUser)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestErr("Invalid Id")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusFound, result)

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "poco a poco bebe")
}
