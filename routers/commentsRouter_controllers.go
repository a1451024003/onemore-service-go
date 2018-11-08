package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["onemore-service-go/controllers:ApplicationDownloadController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:ApplicationDownloadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:ApplicationDownloadController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:ApplicationDownloadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:ApplicationDownloadController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:ApplicationDownloadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:ApplicationDownloadController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:ApplicationDownloadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:ApplicationDownloadController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:ApplicationDownloadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:CodesController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:CodesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:CodesController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:CodesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:CodesController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:CodesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:CodesController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:CodesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:CodesController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:CodesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"],
		beego.ControllerComments{
			Method: "GetNewInfo",
			Router: `/GetNewInfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:HotChcpController"],
		beego.ControllerComments{
			Method: "GetNew",
			Router: `/getNew`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/del_notice`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "GetNotice",
			Router: `/notice`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "DeleteOms",
			Router: `/notice_oms/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:OssController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:OssController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:OssController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:OssController"],
		beego.ControllerComments{
			Method: "GetOss",
			Router: `/oss`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:OssController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:OssController"],
		beego.ControllerComments{
			Method: "GetSts",
			Router: `/sts`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:OtherController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:OtherController"],
		beego.ControllerComments{
			Method: "GetPwd",
			Router: `/encryptPwd`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:PingController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:PingController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:SystemController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:SystemController"],
		beego.ControllerComments{
			Method: "Settings",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:SystemController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:SystemController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:TenantController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:TenantController"],
		beego.ControllerComments{
			Method: "TenantList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:TenantController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:TenantController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:TenantController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:TenantController"],
		beego.ControllerComments{
			Method: "Tenant",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:TenantController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:TenantController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:TenantController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:TenantController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:UploadController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:UploadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:VersionController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:VersionController"],
		beego.ControllerComments{
			Method: "GetNew",
			Router: `/getNew`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:VersionController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:VersionController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:VersionController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:VersionController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:WeatherController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:WeatherController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:WeatherController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:WeatherController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["onemore-service-go/controllers:WeatherController"] = append(beego.GlobalControllerRouter["onemore-service-go/controllers:WeatherController"],
		beego.ControllerComments{
			Method: "City",
			Router: `/city`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
