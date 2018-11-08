package controllers

import (
	"bytes"
	"encoding/base64"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego"
)

// Upload Base64
type UploadController struct {
	beego.Controller
}

func (c *UploadController) URLMapping() {
	c.Mapping("post", c.Post)
}

// @Title 上传bs64
// @Description 上传bs64
// @Param   file     		query    string  	 true       "bs64图片字符串"
// @Success 0 {json} JSONStruct
// @router / [post]
func (c *UploadController) Post() {
	file := c.GetString("file")
	if file == "" {
		c.Data["json"] = JSONStruct{"error", 1001, nil, "上传图片不能为空"}
	} else {
		path := "uploads/"
		filename := GetRandomString(20) + ".png"
		buffer, _ := base64.StdEncoding.DecodeString(file) //成图片文件并把文件写入到buffer
		pic := bytes.NewBuffer(buffer)
		//链接 oss
		client, _ := oss.New(beego.AppConfig.String("OSS_WEB_SERVER"), beego.AppConfig.String("OSS_KEY_ID"), beego.AppConfig.String("OSS_KEY_SECRET"))
		bucket, _ := client.Bucket(beego.AppConfig.String("OSS_BUCKET"))
		timenow := time.Now().Unix()
		datenow := time.Now().Format("2006-01-02")
		tmp_url := path + datenow + "/" + strconv.FormatInt(timenow, 10) + filename
		err := bucket.PutObject(tmp_url, pic) //上传读取的文件
		var url string
		if err == nil {
			url = "http://" + beego.AppConfig.String("OSS_BUCKET") + strings.Replace(beego.AppConfig.String("OSS_WEB_SERVER"), "http://", ".", -1) + "/" + tmp_url
			c.Data["json"] = JSONStruct{"success", 0, url, "上传成功"}
		} else {
			c.Data["json"] = JSONStruct{"error", 1001, err, "上传失败"}
		}
	}

	c.ServeJSON()
}

//生成随机字符串
func GetRandomString(le int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < le; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
