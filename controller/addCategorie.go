package controller

import (
	"auction/model"
	"log"

	"github.com/gin-gonic/gin"
)

func AddCatergory(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)

	var newCatergory *model.Categorie

	err := c.BindJSON(&newCatergory)
	if err != nil {
		log.Println(err.Error())
		return
	}
	errCreateCatergory := db.Table("categories").Create(&newCatergory).Error
	if errCreateCatergory != nil {
		log.Println(errCreateCatergory)
		return
	}
	c.JSON(200, gin.H{
		"add": newCatergory,
	})
}
