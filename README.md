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
<img src="https://scontent.fhan2-2.fna.fbcdn.net/v/t1.0-9/78904067_588444328628402_440389565526573056_n.png?_nc_cat=110&_nc_ohc=j4wvR9ItGnQAQkP_0i7ccBXgJRi-8Q0FbPy5ek2ar1hyCOSA9ueb9Dv_g&_nc_ht=scontent.fhan2-2.fna&oh=9346e89ea633995b39a7ccd8d85f00a9&oe=5E89D543" title="source: imgur.com"  />
<br>
## Cài đặt

Repositories này chứa phần BackEnd của project bao gồm các Rest API.

	```sh 
	//AddAutionHistory
	r.POST("/AddAuctionHistory", controller.AddAuctionHistory)
	//AddCatetories
	r.POST("/AddCategories", controller.AddCatergory)
	//GetCategoriesAll
	r.GET("/GetCategories", controller.GetCategoriesAll)
	//GetAllAuctionHistory
	r.GET("/GetAuctionHistories", controller.GetAuctionHistoryAll)
	//GetAllProduct
	r.GET("/GetProducts", controller.GetAllProduct)
	//GetAllUser
	r.GET("/GetUsers", controller.GetAllUser)
	//AddUser
	r.POST("/UserAdds", controller.AddUser)
	//UpdateUser
	r.POST("/UserUpdates", controller.UpdateUser)
	//DeleteUser
	r.DELETE("/UserDeletes", controller.DeleteUser)
	//AddProduct
	r.POST("/AddProduct", controller.AddProduct)
	//Login
	r.POST("/Login", controller.Login)
	```

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