# fullstack-web-golang-vue-simrs
### Golang Package 
 ```go
 // framework yang digunakan adalah echo
go get github.com/labstack/echo/v4
// untuk keperluan centralized config di project ini
go get github.com/spf13/viper
// untuk keperluan database 
go get gorm.io/gorm
go get gorm.io/driver/postgres
// untuk keperluan routing
go get -u github.com/gorilla/mux
// untuk keperluan logging
go get github.com/sirupsen/logrus
// untuk keperluan api documentation
go get github.com/swaggo/echo-swagger
// untuk validasi
go get github.com/go-playground/validator/v10
// untuk call restapi
go get github.com/valyala/fasthttp
// untuk hot reloa 
go install github.com/cosmtrek/air@latest
 ```