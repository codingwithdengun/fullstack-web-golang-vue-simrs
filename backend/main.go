package main

import (
	"backend/api/configs"
	"fmt"
)

func init() {
	configs.SetupConfiguration()
}
func main() {

	db := configs.InitDB()
	fmt.Println("DATA : ", db)
}
