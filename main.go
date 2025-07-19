package main

import (
	"aplikasi-buku-go/config"
	"aplikasi-buku-go/db"
	"aplikasi-buku-go/routers"
)

func main() {
	config.DbConnection()
	db.DBMigrate(config.DB)
	Init()
}

func Init() {
	defer config.DB.Close()
	r := routers.SetupRoutes()
	r.Run(":8080")
}
