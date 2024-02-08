package handler

import (
	"errors"
	"net/http"

	"github.com/dmuthuraaj/op_zero/usermgt/model"
	"github.com/dmuthuraaj/op_zero/usermgt/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userAdm service.Service
}

func NewUserHandler(uh service.Service) *UserHandler {
	return &UserHandler{
		userAdm: uh,
	}
}

func ErrResponseWithLogger(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"status": code,
		"error":  err.Error(),
		"path":   c.Request.URL.Path,
	})
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var err error
	var user model.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusBadRequest, err)
		return
	}
	err = uh.userAdm.CreateUser(c, &user)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": "user created successfully"})
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := uh.userAdm.GetAllUsers(c)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (uh *UserHandler) GetUserByName(c *gin.Context) {
	// TODO: Change to Id
	userName := c.Param("name")
	if userName == "" {
		ErrResponseWithLogger(c, http.StatusBadRequest, errors.New("name required"))
		return
	}
	user, err := uh.userAdm.GetUserByName(c, userName)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &user})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var err error
	var userUpdate model.User
	err = c.BindJSON(&userUpdate)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusBadRequest, err)
		return
	}
	err = uh.userAdm.UpdateUser(c, &userUpdate)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "user contactInfo updated successfully"})
}

func (uh *UserHandler) DeleteUserByName(c *gin.Context) {
	// TODO: Change to Id
	userName := c.Param("name")
	if userName == "" {
		ErrResponseWithLogger(c, http.StatusBadRequest, errors.New("name required"))
		return
	}
	err := uh.userAdm.DeleteUserByName(c, userName)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"data": "user deleted successfully"})
}
