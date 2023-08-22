package controllers

import (
	"fmt"
	"net/http"
	"playground/config"
	"playground/helpers"
	"playground/models"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	hashedPassword := helpers.HashPassword(user.Password)
	user.Password = hashedPassword

	config.DB.Create(&user)
	fmt.Println(user)
	c.JSON(http.StatusCreated, &user)
}
