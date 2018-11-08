package servers

import (
	"encoding/json"
	"onemore-service-go/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/hprose/hprose-golang/rpc"
	"om-admin-oms/pb/sms"
	"google.golang.org/grpc"
	"fmt"
	"context"
)

type OmsServer struct {
	Send func(phone string, content string, msg string, sendStatus, types int) error
}

var Oms *OmsServer

type SmsServer struct {
}

type Msg struct {
	Mobile string `json:"mobile"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func (c *SmsServer) Init() {
	server := rpc.NewHTTPService()
	server.AddAllMethods(&SmsServer{})
	beego.Handler("/rpc/sms", server)
}

func (c *SmsServer) Send(phone, text string, types int) (interface{}, error) {
	var send Msg
	url := beego.AppConfig.String("SMS_URL")
	apiKey := beego.AppConfig.String("SMS_API_KEY")
	req := httplib.Post(url)
	req.Param("mobile", phone)
	req.Param("text", text)
	req.Param("apikey", apiKey)
	var res interface{}
	err := req.ToJSON(&res)
	r, _ := json.Marshal(res)
	json.Unmarshal([]byte(r), &send)

	serviceAddress := beego.AppConfig.String("OMS_SMS")
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	smsClient := sms.NewSmsLogServiceClient(conn)

	//短信日志详情
	info, _ := smsClient.PostSmsLog(context.Background(), &sms.PostSmsLogParams{Phone:phone, Content:text, Msg:send.Msg, Type:int32(types),SendStatus:int32(send.Code)})

	fmt.Println(string(info.Message))
	//Oms.Send(phone, text, send.Msg, send.Code, types)
	return res, err
}

func (c *SmsServer) Tenant(tenant_id int, kindergarten_id int) error {
	u := models.TenantKindergarten{
		TenantId:       tenant_id,
		KindergartenId: kindergarten_id,
	}
	u.AddTenantKindergarten()

	return nil
}

func (c *SmsServer) Test() string {
	return "sms"
}
