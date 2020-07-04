package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	gethouseinfo "ihome/GetHouseInfo/proto/gethouseinfo"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"strconv"
	"time"
)

type Server struct{}

func (e *Server) GetHouseInfo(ctx context.Context, req *gethouseinfo.Request, rsp *gethouseinfo.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// houseId
	houseId := req.HouseId

	// 用户sessionId： req.SessionId
	sessionId := req.SessionId
	// 连接 redis
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
	// 获取用户id
	userIdTemp := bm.Get(sessionId + "userId")
	userId := int(userIdTemp.([]uint8)[0])

	/* 处理数据 */
	// 从 redis 中获取房屋数据， 没有则去数据库中获取并添加到缓存中
	houseInfo := bm.Get("houseInfo" + houseId)
	if houseInfo != nil {
		// 缓存中有数据
		rsp.UserId = strconv.Itoa(userId)
		rsp.HouseInfo = houseInfo.([]byte)
		return nil
	}

	// 缓存中没有数据， 在数据库中添加
	var house models.House
	o := orm.NewOrm()
	house.Id, _ = strconv.Atoi(houseId)
	_ = o.Read(&house)
	// 关联查询
	o.LoadRelated(&house, "Area")
	o.LoadRelated(&house, "User")
	o.LoadRelated(&house, "Images")
	o.LoadRelated(&house, "Facilities")

	// 将查询到的结果存到缓存中
	houseMarshal, _ := json.Marshal(house)
	_ = bm.Put("houseInfo"+houseId, houseMarshal, time.Second*3600)

	/* 返回数据 */
	rsp.UserId = strconv.Itoa(userId)
	rsp.HouseInfo = houseMarshal

	return nil
}
