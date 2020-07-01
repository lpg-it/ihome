package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"time"

	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	getarea "ihome/GetArea/proto/getarea"
)

type Example struct{}

func (e *Example) GetArea(ctx context.Context, req *getarea.Request, rsp *getarea.Response) error {
	// 初始化 错误码
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	// 1. 从缓存中获取数据
	// 准备连接 redis 信息
	redisConf := map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	}

	// 将 map 转化为 json
	redisConfJson, _ := json.Marshal(redisConf)
	// 创建 redis 句柄
	bm, err := cache.NewCache("redis", string(redisConfJson))
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	// 获取数据
	areaInfo := bm.Get("areaInfo")
	if areaInfo != nil {
		// 缓存中有数据
		var areas []map[string]interface{}
		// 将获取到的数据解码
		json.Unmarshal(areaInfo.([]byte), &areas)

		for _, value := range areas {
			rsp.Data = append(rsp.Data, &getarea.Response_Area{Aid: int32(value["Aid"].(float64)), Aname: value["Aname"].(string)})
		}
		return nil
	}

	// 2. 缓存中没有数据从 mysql 中查找数据
	o := orm.NewOrm()
	var areas []models.Area
	num, err := o.QueryTable("Area").All(&areas)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	if num <= 0 {
		rsp.ErrNo = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 3. 将查找到的数据存到缓存中
	// 将获取到的数据转化为 json 格式
	areasJson, _ := json.Marshal(areas)
	err = bm.Put("areaInfo", areasJson, 3600*time.Second)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 4. 将查找到的数据 按照 proto 的格式 发送给前端
	for _, value := range areas {
		rsp.Data = append(rsp.Data, &getarea.Response_Area{Aid: int32(value.Id), Aname: value.Name})
	}
	return nil
}
