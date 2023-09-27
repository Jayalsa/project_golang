package controllers

import (
	"fmt"
	"jayalsa/project_golang/entities"
	"jayalsa/project_golang/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService interfaces.IUser
}

func InitAuthController(authService interfaces.IUser) *AuthController {
	return &AuthController{AuthService: authService}
}

func (a *AuthController) Register(c *gin.Context) {
	fmt.Println("Invoked controller")
	var user entities.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	result, err := a.AuthService.Register(&user)
	// fmt.Println(result)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}
func (a *AuthController) Login(c *gin.Context) {
	fmt.Println("Invoked controller")
	var user entities.Login
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	result, err := a.AuthService.Login(&user)
	fmt.Println(result)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}
func (a *AuthController) Logout(c *gin.Context) {
	if err := a.AuthService.Logout(); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, "OK")
}
