package servers

import (
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
)

type OssServer struct {
}

func (c *OssServer) Init() {
	server := rpc.NewHTTPService()
	server.AddAllMethods(&OssServer{})
	beego.Handler("/rpc/oss", server)
}

type Oss struct {
	Endpoint  string
	Keyid     string
	Secret    string
	Ossbucket string
}

func (c *OssServer) Upload() Oss {
	var oss Oss

	oss.Endpoint = beego.AppConfig.String("OSS_WEB_SERVER")
	oss.Keyid = beego.AppConfig.String("OSS_KEY_ID")
	oss.Secret = beego.AppConfig.String("OSS_KEY_SECRET")
	oss.Ossbucket = beego.AppConfig.String("OSS_BUCKET")

	return oss
}
