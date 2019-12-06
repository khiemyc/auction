# WebSite Bán Đấu Giá Sản Phẩm 
WebSite cho phép khách hàng tham gia đấu giá các sản phầm như oto, nhà, ..

## Contents
- [Giới thiệu](#Giới-thiệu)
- [Thành viên](#Thành-viên)
- [Công cụ](#Công-cụ)
- [Cấu trúc Database](#Cấu-trúc-Database)
- [Kiến trúc](#Kiến-trúc)
- [Cài đặt](#Cài-đặt)


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

<iframe width="560" height="315" src='https://dbdiagram.io/embed/5de4bc7eedf08a25543e95a9'> </iframe>

## Cài đặt
Repositories này chứa phần BackEnd của project bao gồm các Rest API.

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