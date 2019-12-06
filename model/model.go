package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	jsoniter "github.com/json-iterator/go"
)

// Categorie .
type Categorie struct {
	ID             int    `json:"id"`
	CategoriesName string `json:"NameCategories"`
}

// AuctionHistory .
type AuctionHistory struct {
	ID           int       `json:"id"`
	ProductsID   int       `json:"productId"`
	UsersID      int       `json:"userId"`
	BiddingPrice int       `json:"biddingPrice"`
	TimeToBid    time.Time `json:"timeToBid"`
}

// Product .
type Product struct {
	ID           int
	ProductName  string    `form:"user" json:"ProductName" binding:"required"`
	Description  string    `form:"user" json:"Description" binding:"required"`
	ProductImage string    `form:"user" json:"ProductImage" binding:"required"`
	CategoriesID int       `form:"user" json:"CategoriesID" binding:"required"`
	PriceStart   int       `form:"user" json:"PriceStart" binding:"required"`
	StartTime    time.Time `form:"user" json:"StartTime" binding:"required"`
	EndTime      time.Time `form:"user" json:"EndTime" binding:"required"`
	Status       bool      `form:"user" json:"Status" binding:"required"`
}

// User .
type User struct {
	ID        int    `form:"user" json:"Id" binding:"required"`
	FisrtName string `form:"user" json:"FisrtName" binding:"required"`
	LastName  string `form:"user" json:"LastName" binding:"required"`
	Email     string `form:"user" json:"Email" binding:"required"`
	Wallet    int    `form:"user" json:"Wallet" binding:"required"`
	Phone     string `form:"user" json:"Phone" binding:"required"`
	Role      bool   `form:"user" json:"Role" binding:"required"`
	Password  string `form:"user" json:"Password" binding:"required"`
}

type Config struct {
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		Address  string `json:"address"`
	} `json:"database"`
}

func DecodeDataFromJsonFile(f *os.File, data interface{}) error {
	jsonParser := jsoniter.NewDecoder(f)
	err := jsonParser.Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

func SetupConfig() Config {
	var conf Config

	// Đọc file config.dev.json
	configFile, err := os.Open("config.local.json")
	if err != nil {
		// Nếu không có file config.dev.json thì đọc file config.default.json
		configFile, err = os.Open("config.default.json")
		if err != nil {
			panic(err)
		}
		defer configFile.Close()
	}
	defer configFile.Close()

	// Parse dữ liệu JSON và bind vào conf
	err = DecodeDataFromJsonFile(configFile, &conf)
	if err != nil {
		log.Println("Không đọc được file config.")
		panic(err)
	}

	return conf
}

func ConnectDb(user string, password string, database string, address string) *gorm.DB {
	connectionInfo := fmt.Sprintf(`%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local`, user, password, address, database)

	db, err := gorm.Open("mysql", connectionInfo)
	if err != nil {
		panic(err)
	}
	return db
}
