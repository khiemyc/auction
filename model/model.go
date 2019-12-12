package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	jsoniter "github.com/json-iterator/go"
)

//Login ...
type Login struct {
	Email    string `json:"Email"`
	Password int    `json:"Password"`
}

// Categorie .
type Categorie struct {
	ID             int    `json:"id"`
	CategoriesName string `json:"NameCategories"`
}

// AuctionHistory .
type AuctionHistory struct {
	ID           int    `json:"id"`
	ProductsID   int    `json:"productsId"`
	UsersID      int    `json:"usersId"`
	BiddingPrice int    `json:"biddingPrice"`
	TimeToBid    string `json:"timeToBid"`
}

// Product .
type Product struct {
	ID           int    `json:"id"`
	ProductName  string `json:"productName"`
	Description  string `json:"description"`
	ProductImage string `json:"productImage"`
	CategoriesID int    `json:"Category"`
	PriceStart   int    `json:"priceStart"`
	PriceNow     int    `json:"buyPrice"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	Status       bool   `json:"Sold"`
}

// User .
type User struct {
	ID        int    `json:"Id"`
	FisrtName string `json:"FisrtName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Wallet    int    `json:"Wallet"`
	Phone     string `json:"Phone"`
	Role      bool   `json:"Role"`
	Password  int    `json:"Password"`
}

// userList
type UserList struct {
	FisrtName string `json:"fisrtName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"Phone"`
	Wallet    int    `json:"Wallet"`
	Role      bool   `json:"role"`
	ID        int    `json:"id"`
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
