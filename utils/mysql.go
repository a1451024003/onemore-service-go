package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

func InitMysql() {
	var (
		alias      = beego.AppConfig.String("DB_ALIAS")
		driver     = beego.AppConfig.String("DB_CONNECTION")
		dataSource = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",
			beego.AppConfig.String("DB_USERNAME"),
			beego.AppConfig.String("DB_PASSWORD"),
			beego.AppConfig.String("DB_HOST"),
			beego.AppConfig.String("DB_PORT"),
			beego.AppConfig.String("DB_DATABASE"),
		)
		maxIdle = 1000 //最大空闲
		maxConn = 2000 //最大链接
	)
	orm.RegisterDataBase(alias, driver, dataSource, maxIdle, maxConn)
	orm.DefaultTimeLoc = time.Local

	//日志处理
	filename := time.Now().Format("20060102")
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"storage/logs/%s.log"}`, filename))
	logs.SetLogFuncCall(true) //输出文件名和行号
	logs.Async()              //异步输出
}
