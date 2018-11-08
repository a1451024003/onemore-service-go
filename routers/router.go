// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"onemore-service-go/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//引入
	pingCtl := &controllers.PingController{}
	errorCtl := &controllers.ErrorController{}

	ns := beego.NewNamespace("/api/v2/onemore",

		beego.NSNamespace("/oss",
			beego.NSInclude(
				&controllers.OssController{},
			),
		),
		beego.NSNamespace("/app/upload",
			beego.NSInclude(
				&controllers.UploadController{},
			),
		),

		beego.NSNamespace("/ping",
			beego.NSInclude(
				pingCtl,
			),
		),

		beego.NSNamespace("/tenant",
			beego.NSInclude(
				&controllers.TenantController{},
			),
		),

		beego.NSNamespace("/other",
			beego.NSInclude(
				&controllers.OtherController{},
			),
		),

		beego.NSNamespace("/codes",
			beego.NSInclude(
				&controllers.CodesController{},
			),
		),

		beego.NSNamespace("/chcp",
			beego.NSInclude(
				&controllers.HotChcpController{},
			),
		),
		beego.NSNamespace("/version",
			beego.NSInclude(
				&controllers.VersionController{},
			),
		),
		beego.NSNamespace("/notice",
			beego.NSInclude(
				&controllers.NoticeController{},
			),
		),

		beego.NSNamespace("/conf",
			beego.NSInclude(
				&controllers.SystemController{},
			),
		),
		beego.NSNamespace("/application_download",
			beego.NSInclude(
				&controllers.ApplicationDownloadController{},
			),
		),
		beego.NSRouter("/system", &controllers.NoticeController{}, "post:PostSystem"),
		beego.NSNamespace("/weather",
			beego.NSInclude(
				&controllers.WeatherController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//错误处理
	beego.ErrorController(errorCtl)
}
