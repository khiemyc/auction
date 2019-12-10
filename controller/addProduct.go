package controller

import (
	"auction/model"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)

	var Products model.Product

	err := c.BindJSON(&Products)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		Products.ProductImage = filename
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	errCreateProduct := db.Table("products").Create(Products).Error
	if errCreateProduct != nil {
		log.Println(errCreateProduct)
		return
	}
	c.JSON(200, gin.H{
		"addProduct": Products,
	})
}
