package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	getindex "ihome/GetIndex/proto/getindex"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"time"
)

type Server struct{}

func (e *Server) GetIndex(ctx context.Context, req *getindex.Request, rsp *getindex.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 先从缓存中获取， 没有则从数据库获取并加入缓存
	redisConfig, _ := json.Marshal(map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	})
	bm, err := cache.NewCache("redis", string(redisConfig))
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	homePageData := bm.Get("homePageData")
	if homePageData != nil {
		// 缓存中有数据, 直接将二进制数据发送给客户端
		rsp.Data = homePageData.([]byte)
		return nil
	}

	// 缓存中没有数据， 从数据库中查询
	o := orm.NewOrm()
	var houses []models.House
	_, err = o.QueryTable("House").Limit(models.HomePageMaxHouses).All(&houses)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	/* 处理数据 */
	var data []interface{}
	for _, house := range houses {
		o.LoadRelated(&house, "Area")
		o.LoadRelated(&house, "User")
		o.LoadRelated(&house, "Images")
		o.LoadRelated(&house, "Facilities")

		data = append(data, house.ToHouseInfo())
	}

	homePageData, _ = json.Marshal(data)
	bm.Put("homePageData", homePageData, time.Second*3600)

	/* 返回数据 */
	rsp.Data = homePageData.([]byte)

	return nil
}
