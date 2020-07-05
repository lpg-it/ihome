package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"strconv"
	"time"

	posthouses "ihome/PostHouses/proto/posthouses"
)

type Server struct{}

func (e *Server) PostHouses(ctx context.Context, req *posthouses.Request, rsp *posthouses.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 用户 sessionId：req.SessionId

	// 连接 redis, 获取到 用户id
	redisConfig, _ := json.Marshal(map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	})
	bm, err := cache.NewCache("redis", string(redisConfig))
	if err != nil {
		// 连接 redis 失败
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	userIdTemp := bm.Get(req.SessionId + "userId")
	userId := int(userIdTemp.([]uint8)[0])

	/* 处理数据 */
	o := orm.NewOrm()
	houseInfo := make(map[string]interface{})
	var house models.House

	_ = json.Unmarshal(req.Data, &houseInfo)
	house.Title = houseInfo["title"].(string)
	house.Price, _ = strconv.Atoi(houseInfo["price"].(string))
	house.Address = houseInfo["address"].(string)
	house.RoomCount, _ = strconv.Atoi(houseInfo["room_count"].(string))
	house.Acreage, _ = strconv.Atoi(houseInfo["acreage"].(string))
	house.Unit = houseInfo["unit"].(string)
	house.Capacity, _ = strconv.Atoi(houseInfo["capacity"].(string))
	house.Beds = houseInfo["beds"].(string)
	house.Deposit, _ = strconv.Atoi(houseInfo["deposit"].(string))
	house.MinDays, _ = strconv.Atoi(houseInfo["min_days"].(string))
	house.MaxDays, _ = strconv.Atoi(houseInfo["max_days"].(string))
	// 地区
	areaId, _ := strconv.Atoi(houseInfo["area_id"].(string))
	var area models.Area
	area.Id = areaId
	_ = o.Read(&area)
	house.Area = &area

	// 设施
	var facilities []*models.Facility
	for _, value := range houseInfo["facility"].([]interface{}) {
		facilityId, _ := strconv.Atoi(value.(string))
		var facility models.Facility
		facility.Id = facilityId
		_ = o.Read(&facility)
		facilities = append(facilities, &facility)
	}
	house.Facilities = facilities

	// 添加信息
	var user models.User
	user.Id = userId
	_ = o.Read(&user)
	house.User = &user
	_, err = o.Insert(&house)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 插入到房源与设施信息的多对多表中
	num, err := o.QueryM2M(&house, "Facilities").Add(facilities)
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

	// 设置 session
	bm.Put(req.SessionId + "name", user.Name, time.Second * 3600)
	bm.Put(req.SessionId + "userId", string(user.Id), time.Second * 3600)
	bm.Put(req.SessionId + "mobile", user.Mobile, time.Second * 3600)

	/* 返回数据 */
	rsp.HouseId = strconv.Itoa(house.Id)

	return nil
}
