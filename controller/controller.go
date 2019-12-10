package controller

import (
	"auction/model"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// func RegisterFormAddUser(c *gin.Context) {
// 	c.HTML(http.StatusOK, "addUsers.html", gin.H{})
// }

// func RegisterFormAddCatergory(c *gin.Context) {
// 	c.HTML(http.StatusOK, "addCatergorys.html", gin.H{})
// }
/*

/******SINGLETON Database Connection****/
var once sync.Once

//DatabaseB ...It hold the pointer to database.
type DatabaseB struct {
	Db *gorm.DB
}

//variance global
var instance *DatabaseB

//GetDBInstance ...Use this function go fetch database instance.
func GetDBInstance() *DatabaseB {
	once.Do(func() { //do not allow repeating
		//thread safe
		instance = &DatabaseB{}
	})

	return instance
}

//GetCategoriesAll ...Get all
// @Description Get tất cả dữ liệu ở bảng Categories trả về  Json
// @Accept json
// @Success 200 {object} model.Categorie
// @Router /GetAllCategories [get]
func GetCategoriesAll(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)
	var catergorysAll []model.Categorie
	errGetUser := db.Find(&catergorysAll).Error
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}
	for _, v := range catergorysAll {
		log.Println("Employee", v)
	}
	c.JSON(200, catergorysAll)
}

//GetAuctionHistoryAll ...GET all
//AuctionHistoryAll ...Get all
// @Description Get tất cả dữ liệu ở bảng AuctionHistoryAll trả về  Json
// @Accept json
// @Success 200 {object} model.AuctionHistory
// @Router /GetAllAuctionHistory [get]
func GetAuctionHistoryAll(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)
	var auctionHistoryAll []model.AuctionHistory
	errGetUser := db.Table("auction_history").Model(&model.AuctionHistory{}).Select("*").Scan(&auctionHistoryAll).Error
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}
	c.JSON(200, auctionHistoryAll)
}

//GetProductAll ...Get all
//GetProductAll ...Get all
// @Description Get tất cả dữ liệu ở bảng GetProductAll trả về  Json
// @Accept json
// @Success 200 {object} model.Product
// @Router /GetAllProduct [get]
func GetProductAll(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)
	var GetProductAll []model.Product
	errGetProduct := db.Table("products").Model(&model.Product{}).Select("*").Scan(&GetProductAll).Error
	if errGetProduct != nil {
		log.Println(errGetProduct)
		return
	}
	c.JSON(200, GetProductAll)
}

//GetUserAll ...Get all
// @Description Get tất cả dữ liệu ở bảng GetUserAll trả về  Json
// @Accept json
// @Success 200 {object} model.User
// @Router /GetAllUser [get]
func GetUserAll(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)
	var getUserAll []model.User
	errGetUser := db.Table("users").Model(&model.User{}).Select("*").Scan(&getUserAll).Error
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}
	c.JSON(200, getUserAll)
}

type LoginForm struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func EmxampleAPI(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var loginForm *LoginForm
	err := c.ShouldBindJSON(&loginForm)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("aa %+v", loginForm)
	c.JSON(200, gin.H{
		"test": loginForm,
	})
}

func AddUser(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)

	var addUser *model.User

	err := c.ShouldBindJSON(&addUser)
	if err != nil {
		log.Println(err.Error())
		return
	}
	errCreateUser := db.Table("users").Create(&addUser).Error
	if errCreateUser != nil {
		log.Println(errCreateUser)
		return
	}

	c.JSON(200, gin.H{
		"add": addUser,
	})
}

func UpdateUser(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)

	var addUser *model.User

	err := c.ShouldBindJSON(&addUser)
	if err != nil {
		log.Println(err.Error())
		return
	}
	errUpdateUser := db.Table("users").Save(&addUser).Error
	if errUpdateUser != nil {
		log.Println(errUpdateUser)
		return
	}

	c.JSON(200, gin.H{
		"update": addUser,
	})
}

func DeleteUser(c *gin.Context) {
	db := GetDBInstance().Db
	db.LogMode(true)

	var addUser *model.User

	err := c.ShouldBindJSON(&addUser)
	if err != nil {
		log.Println(err.Error())
		return
	}
	errDeleteUser := db.Table("users").Where("id = ?", addUser.ID).Delete(&addUser.ID).Error
	if errDeleteUser != nil {
		log.Println(errDeleteUser)
		return
	}

	c.JSON(200, gin.H{
		"delete": addUser,
	})
}

// func AddUser(c *gin.Context) {
// 	config := model.SetupConfig()
// 	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
// 	defer db.Close()
// 	db.LogMode(true)

// 	var newUser = model.User{
// 		Full_name: c.PostForm("full_name"),
// 		Email:     c.PostForm("email"),
// 		Password:  c.PostForm("password"),
// 		Info:      c.PostForm("info"),
// 	}

// 	errCreateUser := db.Create(&newUser).Error
// 	if errCreateUser != nil {
// 		log.Println(errCreateUser)
// 		return
// 	}
// }


func GetNameCategories(c *gin.Context) {
	// config := model.SetupConfig()
	// db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	// defer db.Close()

	db := GetDBInstance().Db

	db.LogMode(true)
	var GetNames []model.Categorie
	name := c.Query("name")

	errGetNamCategory := db.Table("categories").Select("id, categories_name").
		Where("categories_name = ?", name).Scan(&GetNames).Error
	if errGetNamCategory != nil {
		log.Println(errGetNamCategory)
		return
	}
	c.JSON(200, GetNames)
}

func GetUserBy(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var userInfo []model.User

	id := c.Query("id")
	name := c.Query("name")
	email := c.Query("email")
	info := c.Query("info")
	waller := c.Query("waller")

	errGetUser := db.Table("users").Select("id, full_name, email,password,info,role,waller").
		Where("id = ?", id).Or("full_name = ?", name).Or("email = ?", email).Or("info = ?", info).Or("waller = ?", waller).Scan(&userInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}

	log.Println("Employee", userInfo)
	c.JSON(200, userInfo)
}

func GetAllUser(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var UserAll []model.User
	errGetUser := db.Find(&UserAll).Error
	if errGetUser != nil {
		log.Println(errGetUser)
		return
	}

	for _, v := range UserAll {
		log.Println("Employee", v)
	}

	c.JSON(200, UserAll)
}

func GetAllProduct(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var ProductAll []model.Product
	errGetProduct := db.Find(&ProductAll).Error
	if errGetProduct != nil {
		log.Println(errGetProduct)
		return
	}
	c.JSON(200, ProductAll)
}
