# WebSite Bán Đấu Giá Sản Phẩm 
WebSite cho phép khách hàng tham gia đấu giá các sản phầm như oto, nhà, ..

## Contents
- [Giới thiệu](#Giới-thiệu)
- [Thành viên](#Thành-viên)
- [Công cụ](#Công-cụ)

## Giới thiệu
- Người quản trị đăng các sản phẩm lên web, mở các phiên đấu giá
- Khách hàng sẽ tham gia trực tiếp đấu giá tại bất gì đâu

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

<br>
<img src="https://scontent.fhan3-1.fna.fbcdn.net/v/t1.0-9/78904912_588274935312008_6014814570894327808_n.png?_nc_cat=110&_nc_ohc=XkvlE0pW-i8AQk_yKXvFU_COCQUgNzZLU_yM0KyTJuG4Ui_wl4Lz9u3uA&_nc_ht=scontent.fhan3-1.fna&oh=f1cf8a86bfa0241dc8e923875f1d5510&oe=5E6B245E" title="source: imgur.com"  />
<br>
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