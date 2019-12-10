package controller

import (
	"auction/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)

	var Login model.User
	if err := c.ShouldBindJSON(&Login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// errLogin := db.Exec("SELECT * FROM `users` WHERE email = ? AND password= ?", Login.Email, Login.Password).Error
	errLogin := db.Where("email = ? AND password=?", Login.Email, Login.Password).Find(&Login).Error

	if errLogin != nil {
		log.Println(errLogin)
		c.JSON(404, gin.H{
			"err": "sai tai khoan mat khau",
		})
		return
	}
	c.JSON(200, Login)
}
