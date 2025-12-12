package main

import (
	"github.com/zhang/LibraryMS/internal/bootstrap"
	"github.com/zhang/LibraryMS/internal/config"
	"github.com/zhang/LibraryMS/internal/model"
)

func main() {
	config.InitConfig()
	bootstrap.InitDB()

	bootstrap.DB.AutoMigrate(model.User{})

}
