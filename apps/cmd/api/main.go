package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	v1 "github.com/zhang/LibraryMS/internal/api/v1"
	"github.com/zhang/LibraryMS/internal/bootstrap"
	"github.com/zhang/LibraryMS/internal/config"
	"github.com/zhang/LibraryMS/internal/repository"
	"github.com/zhang/LibraryMS/internal/router"
	"github.com/zhang/LibraryMS/internal/service"
)

func main() {
	// init config
	config.InitConfig()
	fmt.Println(config.Cfg)
	// connection database
	bootstrap.InitDB()

	// load gin
	r := gin.Default()

	// load router
	user_repo := repository.NewIUserRepository(bootstrap.DB)
	user_service := service.NewIUserService(user_repo)
	user_ctl := v1.NewUserController(user_service)
	router.UserRouter(r, user_ctl)

	adds := fmt.Sprintf("%s:%s", config.Cfg.Server.Host, config.Cfg.Server.Post)
	if err := r.Run(adds); err != nil {
		panic("gin 服务启动失败")
	}

}
