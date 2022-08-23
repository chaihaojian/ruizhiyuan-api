package main

import (
	"fmt"
	"go.uber.org/zap"
	"ruizhiyuan/config"
	"ruizhiyuan/dao/mysql"
	"ruizhiyuan/logger"
	"ruizhiyuan/pkg/snowflake"
	"ruizhiyuan/routes"
)

func main() {
	//加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("Init config faild, err:%v\n", err)
	}
	//初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("Init logger faild, err:%v\n", err)
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//初始化MySQL连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("Init mysql faild, err:%v\n", err)
	}
	defer mysql.Close()
	//初始化ID生成器
	if err := snowflake.Init(); err != nil {
		fmt.Printf("Init snowflake faild, err:%v\n", err)
	}
	//注册路由
	r := routes.Setup()

	r.Run()
}
