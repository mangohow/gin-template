package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mangohow/easygin"
	"github.com/mangohow/gin-template/conf"
	"github.com/mangohow/gin-template/internal/dao/db"
	"github.com/mangohow/gin-template/internal/routes"
	"github.com/mangohow/gin-template/pkg/logger"
)

func main() {
	// 初始化配置
	if err := conf.LoadConf(); err != nil {
		panic(fmt.Errorf("load conf failed, reason:%s", err.Error()))
	}

	// 初始化日志
	if err := logger.InitLogger(); err != nil {
		panic(fmt.Errorf("init logger error, reason:%v", err))
	}

	// 初始化数据库
	if err := db.InitMysql(); err != nil {
		// panic(fmt.Errorf("init mysql failed, reason:%s", err.Error()))
	}

	// 创建gin路由
	easyGin := easygin.NewWithEngine(gin.Default())
	easygin.SetLogOutput(logger.Logger().Out)

	// 注册路由
	routes.Register(easyGin)

	easyGin.SetAfterCloseHandlers(func() {
		db.CloseMysql()
		fmt.Println("close mysql connection.")
	})

	err := easyGin.ListenAndServe(fmt.Sprintf("%s:%d", conf.Config().Server.Host, conf.Config().Server.Port))
	if err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("server closed")
		}

		fmt.Println(err)
	}
}
