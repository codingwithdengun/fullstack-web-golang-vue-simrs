package main

import (
	"backend/api/configs"
)

func init() {
	configs.SetupConfiguration()
}
func main() {
	configs.InitDB()
}
