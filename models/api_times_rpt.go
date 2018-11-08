package models

type ApiTimesRpt struct {
	Date        string `orm:"column(date);size(45);null" description:"日期"`
	ServiceJson string `orm:"column(service_json);size(455);null" description:"json格式的服务名称,次数"`
}
