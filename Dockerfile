FROM golang:alpine
RUN mkdir $GOPATH/src/auction
ADD . $GOPATH/src/auction
WORKDIR $GOPATH/src/auction
RUN apk update && apk add git && go get -u github.com/gin-gonic/gin github.com/jinzhu/gorm github.com/go-sql-driver/mysql github.com/json-iterator/go && go build -o main 
ENTRYPOINT ["./main"]