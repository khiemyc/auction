# WebSite Bán Đấu Giá Sản Phẩm 
WebSite cho phép khách hàng tham gia đấu giá các sản phầm như oto, nhà, ..

## Contents
- [Giới thiệu](#Giới-thiệu)
- [Thành viên](#Thành-viên)
- [Công cụ](#Công-cụ)

## Giới thiệu
Website cho phép người quản trị đăng các sản phẩm lên web, mở các phiên đấu giá
Khách hàng sẽ tham gia trực tiếp đấu giá tại bất gì đâu

## Thành viên
- BackEnd:
- FronEnd:

## Công cụ
Project sử dụng các công cụ:
- Virtualization: docker
- Version control: git
- Database: mysql, phpmyadmin
- Golang
- Gin-gonic : https://github.com/gin-gonic
- Gorm : https://github.com/jinzhu/gorm
- IDE : Visual Studio Code
- Document Api: GoSwagger
- Proxy: Localxpose: https://localxpose.io/

## Cấu trúc Database
Cấu trúc các bảng trong database của project

https://dbdiagram.io/d/5de4bc7eedf08a25543e95a9

## Cài đặt

Repositories này chứa phần BackEnd của project bao gồm các Rest API.

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

2. Cài đặt các thư viện cần thiết
```sh
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/gin-gonic/contrib/jwt
$ go get -u github.com/swaggo/files
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/jinzhu/gorm
$ go get -u github.com/jinzhu/gorm/dialects/
$ go get -u github.com/json-iterator/go
$ go get -u golang.org/x/crypto/bcrypt
```