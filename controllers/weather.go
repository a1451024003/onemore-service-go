package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"onemore-service-go/models"
	"github.com/astaxie/beego/validation"
	"time"
	"github.com/astaxie/beego/orm"
)

type WeatherController struct {
	beego.Controller
}

func (c *WeatherController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("City", c.City)
}

type WeatherData struct {
	Province string `json:"province"`
	CityName string `json:"city_name"`
	DayWeather string `json:"day_w"`
	DayTemp string `json:"day_T"`
	NightWeather string `json:"night_w"`
	NightTemp string `json:"night_T"`
	DayWind string `json:"day_wind"`
	DayWindScale string `json:"day_speed"`
	NightWind string `json:"night_wind"`
	NightWindScale string `json:"night_speed"`
	AirQuality string `json:"air"`
	WeatherList []string `json:"weather_list"`
	TempList []string `json:"T_list"`
	WindList []string `json:"windd_list"`
	TimeList []string `json:"time_list"`
	HumidityList []string `json:"xdsd_list"`
}

// @Title 保存天气
// @Description 保存天气
// @Param   data     		formData    json  		true         "天气数据"
// @Success 0 {json} JSONStruct
// @Failure 1003 保存失败
// @router / [post]
func (c *WeatherController) Post() {
	data := c.GetString("data")
	if data == "" {
		c.Data["json"] = JSONStruct{"error", 1001, "", "参数不能为空"}
		c.ServeJSON()
		c.StopRun()
	}

	var wds []WeatherData
	err := json.Unmarshal([]byte(data), &wds)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1001, "", "数据解析失败"}
		c.ServeJSON()
		c.StopRun()
	}

	var w models.Weather
	o := orm.NewOrm()
	o.Begin()
	if err := w.Clear(o); err != nil {
		o.Rollback()

		c.Data["json"] = JSONStruct{"error", 1004, "", "清理weather表失败"}
		c.ServeJSON()
		c.StopRun()
	}
	var ws []models.Weather
	count := 0
	for _, val := range wds {
		for key, v := range val.WeatherList {
			wtt := models.Weather{
				Province:val.Province,
				CityName:val.CityName,
				Date:time.Now().Format("2006-01-02"),
				DayWeather:val.DayWeather,
				NightWeather:val.NightWeather,
				DayTemp:val.DayTemp,
				NightTemp:val.NightTemp,
				DayWind:val.DayWind,
				DayWindScale:val.DayWindScale,
				NightWind:val.NightWind,
				NightWindScale:val.NightWindScale,
				AirQuality:val.AirQuality,
				Time:val.TimeList[key],
				Weather:v,
				Temp:val.TempList[key],
				Humidity:val.HumidityList[key],
				Wind:val.WindList[key],
			}
			ws = append(ws, wtt)
			count++
			if count % 1000 == 0 {
				if err := w.Save(ws, o); err != nil {
					o.Rollback()

					c.Data["json"] = JSONStruct{"error", 1003, "", "保存失败"}
					c.ServeJSON()
					c.StopRun()
				}
				ws = make([]models.Weather, 0)
			}
		}
	}
	if err := w.Save(ws, o); err == nil {
		o.Commit()

		c.Data["json"] = JSONStruct{"success", 0, "", "保存成功"}
	} else {
		o.Rollback()

		c.Data["json"] = JSONStruct{"error", 1003, "", "保存失败"}
	}
	c.ServeJSON()
}

// @Title 获取最新天气
// @Description 获取最新天气
// @Param   city_name     		query    string  	true        "城市名称"
// @Success 0 {json} JSONStruct
// @Failure 1002 当前城市天气无数据
// @Failure 1005 获取失败
// @router / [get]
func (c *WeatherController) Get() {
	province := c.GetString("province")
	cityName := c.GetString("city_name")

	valid := validation.Validation{}
	valid.Required(province, "province").Message("省份名称不能为空")
	valid.Required(cityName, "city_name").Message("城市名称不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, "", valid.Errors[0].Message}

		c.ServeJSON()
		c.StopRun()
	}

	w := models.Weather{Province:province, CityName:cityName}

	if err := w.Get(); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, w, "获取成功"}
	} else if err == orm.ErrNoRows {
		c.Data["json"] = JSONStruct{"error", 1002, "", "当前城市天气无数据"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, "", "获取失败"}
	}

	c.ServeJSON()
}

// @Title 获取城市列表
// @Description 获取城市列表
// @Param   parent_id     		query    int  	false        "父级ID，0为省份，大于0为市县"
// @Success 0 {json} JSONStruct
// @Failure 1002 父级ID不存在
// @Failure 1005 获取失败
// @router /city [get]
func (c *WeatherController) City() {
	parentId, _ := c.GetInt("parent_id", 0)

	ct := models.City{ParentId:parentId}

	if res, num, err := ct.Get(); err == nil && num > 0 {
		c.Data["json"] = JSONStruct{"success", 0, res, "获取成功"}
	} else if err == nil && num <= 0 {
		c.Data["json"] = JSONStruct{"error", 1002, "", "父级ID不存在"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, "", "获取失败"}
	}

	c.ServeJSON()
}
