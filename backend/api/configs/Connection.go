package configs

import (
	"fmt"
	"net/url"
)

func initDB() *gorm.DB {

	val := url.Values{}

	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%v)/%s?%s`, configs.)
}
