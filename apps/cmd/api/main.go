package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zhang/LibraryMS/internal/bootstrap"
	"github.com/zhang/LibraryMS/internal/config"
)

func main() {
	// init config
	config.InitConfig()
	fmt.Println(config.Cfg)
	// connection database
	bootstrap.InitDB()

	// load gin
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		c.String(200, "hello gin")
	})

	adds := fmt.Sprintf("%s:%s", config.Cfg.Server.Host, config.Cfg.Server.Post)
	if err := r.Run(adds); err != nil {
		panic("gin 服务启动失败")
	}

}
