package controller

import (
	"go_test/config"
	"go_test/models"
	"net/http"

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
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(http.StatusOK, &user)

}
