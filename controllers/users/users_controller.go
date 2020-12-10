package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nelbermora/bookstore_users-api/domain/users"
	"github.com/nelbermora/bookstore_users-api/services"
	"github.com/nelbermora/bookstore_users-api/utils/errors"
)

// CreateUser calls Service Layer for Create Users
func CreateUser(c *gin.Context) {
	var newUser users.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		restErr := errors.NewBadRequestErr("Invalid Json")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	result, saveErr := services.UsersService.CreateUser(newUser)
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
	result, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	isPublicReq := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusFound, result.Marshall(isPublicReq))

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "poco a poco bebe")
}

func UpdateUser(c *gin.Context) {
	var newUser users.User
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestErr("Invalid Id")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		restErr := errors.NewBadRequestErr("Invalid Json")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	newUser.Id = userId
	result, updErr := services.UsersService.Update(newUser)
	if updErr != nil {
		c.JSON(updErr.Code, updErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Delete(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestErr("Invalid Id")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	delErr := services.UsersService.DeleteUser(userId)
	if delErr != nil {
		c.JSON(delErr.Code, delErr)
		return
	}
	c.JSON(200, map[string]string{"status": "deleted"})

}

func FindByStatus(c *gin.Context) {
	status := c.Query("status")
	result, err := services.UsersService.FindByStatus(status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result)

}
