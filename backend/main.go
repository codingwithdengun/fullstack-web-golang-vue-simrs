package main

func init() {
	configs.SetupConfiguration()
}
func main() {

	db := configs.InitDB()
	route.CreateHandler(db)
}
