package main

import (
	_ "onemore-service-go/routers"

	"onemore-service-go/servers"
	"onemore-service-go/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func init() {
	utils.InitMysql()
	redisConf := fmt.Sprintf("%s:%s", beego.AppConfig.String("REDIS_HOST"), beego.AppConfig.String("REDIS_PORT"))
	utils.InitRedix(redisConf, "", 200, 60)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}
	smsServer := servers.SmsServer{}
	smsServer.Init()
	ossServer := servers.OssServer{}
	ossServer.Init()
	NoticeServer := servers.NoticeServer{}
	NoticeServer.Init()
	beego.Run()
}
