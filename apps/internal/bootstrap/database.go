package bootstrap

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zhang/LibraryMS/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	// 读取 config.yaml 中的数据库配置
	cfg := config.Cfg.DataSource

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	fmt.Println(dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Millisecond * 200,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err)
		panic("数据库连接失败")
	}

	DB = db

	log.Println("数据库初始化完毕")
}
