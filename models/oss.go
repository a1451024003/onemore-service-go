package models

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"hash"
	"io"
	"time"

	"github.com/astaxie/beego"
)

var accessKeyId = beego.AppConfig.String("OSS_KEY_ID")
var accessKeySecret = beego.AppConfig.String("OSS_KEY_SECRET")
var host = beego.AppConfig.String("OSS_RESOURCES_SERVE")
var expireTime, _ = beego.AppConfig.Int64("EXPIRE_TIME")
var androidAccessKeyId = beego.AppConfig.String("OSS_ANDROID_ACCESSKEYID")
var androidAccessKeySecret = beego.AppConfig.String("OSS_ANDROID_ACCESSKEYSECRET")
var androidRoleArn = beego.AppConfig.String("OSS_ANDROID_ROLEARN")
var androidRolesessionname = beego.AppConfig.String("OSS_ANDROID_ROLESESSIONNAME")
var androidDurationseconds, _ = beego.AppConfig.Int("OSS_ANDROID_DURATIONSECONDS")

func getGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
}

type SecurityToken struct {
	AccessKeyID string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	SecurityToken string `json:"security_token"`
	Expiration time.Time `json:"expiration"`
} 

func GetPolicyToken(path string) PolicyToken {
	now := time.Now().Unix()
	expireEnd := now + expireTime
	var tokenExpire = getGmtIso8601(expireEnd)
	if path == "" {
		path = "images"
	}
	env := beego.AppConfig.String("OSS_WEB_BUCKET")
	uploadDir := path + "/" + env + "/" + time.Now().Format("20060102")

	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, uploadDir)
	config.Conditions = append(config.Conditions, condition)

	//calucate signature
	result, _ := json.Marshal(config)
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(accessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var policyToken PolicyToken
	policyToken.AccessKeyId = accessKeyId
	policyToken.Host = host
	policyToken.Expire = expireEnd
	policyToken.Signature = string(signedStr)
	policyToken.Directory = uploadDir
	policyToken.Policy = string(debyte)

	return policyToken
}

func GetSecurityToken() (map[string]interface{}, error) {
	stsClient := NewClient(androidAccessKeyId, androidAccessKeySecret, androidRoleArn, androidRolesessionname)

	var  v map[string]interface{}
	resp, err := stsClient.AssumeRole(uint(androidDurationseconds))
	if err != nil {
		return v, err
	}

	securityToken := SecurityToken{
		resp.Credentials.AccessKeyId,
		resp.Credentials.AccessKeySecret,
		resp.Credentials.SecurityToken,
		resp.Credentials.Expiration,
	}

	aa ,_ := json.Marshal(securityToken)
	json.Unmarshal(aa,&v)
	v["endpoint"] =  beego.AppConfig.String("OSS_PICTUREBOOK_SERVER")
	v["oss_bucket"] = beego.AppConfig.String("OSS_BUCKET_PICTUREBOOK")

	return v, nil
}
