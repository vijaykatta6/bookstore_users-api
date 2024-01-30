package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" // <-- Updated import path
	users_domain "github.com/vijaykatta6/bookstore-users-api/Domain/users"
	"github.com/vijaykatta6/bookstore-users-api/services"
	rest_errors "github.com/vijaykatta6/bookstore-users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := rest_errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUSer(userId)
	if getErr != nil {
		// handle the error
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users_domain.User
	fmt.Println(user)

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// handle the error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	fmt.Println(result)

	c.JSON(http.StatusCreated, result)
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "Implement me !!!")
// }
