package main

import (
	"github.com/CherkashinAV/finance_app/initializers"
	"github.com/CherkashinAV/finance_app/models"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}