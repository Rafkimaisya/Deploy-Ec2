package main

import (
	"fmt"

	"github.com/jackthepanda96/Belajar-Rest.git/config"
	"github.com/jackthepanda96/Belajar-Rest.git/infrastructure/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	println("Test Server")
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()

	fmt.Println("Menjalankan program ....")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))

}
