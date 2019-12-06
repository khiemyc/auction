package main

import (
	"auction/controller"
	"auction/model"

	"github.com/gin-gonic/gin"

	_ "auction/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	//Reference to singleton variance
	databaseB := controller.GetDBInstance()
	//Open Database from JSON config
	config := model.SetupConfig()
	databaseB.Db = model.ConnectDb(config.Database.User,
		config.Database.Password,
		config.Database.Database,
		config.Database.Address)
	defer databaseB.Db.Close()
	databaseB.Db.LogMode(true)

	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8083/auction/doc.json") // The url pointing to API definition
	r.GET("/auction/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	
	//GetCategoriesAll
	r.GET("/GetAllCategories", controller.GetCategoriesAll)
	//GetAllAuctionHistory
	r.GET("/GetAllAuctionHistory", controller.GetAuctionHistoryAll)
	//GetAllProduct
	r.GET("/GetAllProduct", controller.GetAllProduct)
	//GetAllUser
	r.GET("/GetAllUser", controller.GetAllUser)
	//AddUser
	r.POST("/UserAdd", controller.AddUser)
	//UpdateUser
	r.POST("/UserUpdate", controller.UpdateUser)
	//DeleteUser
	r.POST("/UserDelete", controller.DeleteUser)

	r.POST("/AddCatergory", controller.AddCatergory)
	// test
	r.POST("/DeleteUser", controller.DeleteUser)
	r.Run(":8083")
}
