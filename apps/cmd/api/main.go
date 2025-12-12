package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zhang/LibraryMS/internal/bootstrap"
	"github.com/zhang/LibraryMS/internal/config"
	v1 "github.com/zhang/LibraryMS/internal/handler/v1"
	"github.com/zhang/LibraryMS/internal/repository"
	"github.com/zhang/LibraryMS/internal/router"
	"github.com/zhang/LibraryMS/internal/service"
)

func main() {
	// 1. 初始化配置
	config.InitConfig()
	log.Printf("配置加载完成: %+v", config.Cfg)

	// 2. 初始化数据库连接
	bootstrap.InitDB()

	// 3. 设置 Gin 模式
	gin.SetMode(config.Cfg.Server.Mode)

	// 4. 创建 Gin 引擎
	r := gin.Default()

	// 5. 依赖注入：初始化各层
	// User module
	userRepo := repository.NewIUserRepository(bootstrap.DB)
	userService := service.NewIUserService(userRepo)
	userHandler := v1.NewUserHandler(userService)
	// Book module
	bookRepo := repository.NewIBookRepository(bootstrap.DB)
	bookService := service.NewIBookService(bookRepo)
	bookHandler := v1.NewBookHandler(bookService)
	// 6. 注册路由
	router.SetupRouter(r,
		userHandler,
		bookHandler,
	)
	// 7. 启动服务
	addr := fmt.Sprintf("%s:%s", config.Cfg.Server.Host, config.Cfg.Server.Port)
	log.Printf("服务启动: http://%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
