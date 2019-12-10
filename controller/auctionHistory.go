package controller

import (
	"auction/model"
	"log"

	"github.com/gin-gonic/gin"
)

func AddAuctionHistory(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)

	var newAuctionHistory *model.AuctionHistory

	err := c.ShouldBindJSON(&newAuctionHistory)

	if err != nil {
		log.Println(err.Error())
		return
	}

	errCreateUser := db.Table("auction_history").Create(&newAuctionHistory).Error
	if errCreateUser != nil {
		log.Println(errCreateUser)
		return
	}

	c.JSON(200, gin.H{
		"add": newAuctionHistory,
	})
}
