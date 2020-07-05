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

	getuserorder "ihome/GetUserOrder/proto/getuserorder"
)

type Server struct{}

func (e *Server) GetUserOrder(ctx context.Context, req *getuserorder.Request, rsp *getuserorder.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 获取 sessionId 查找对应用户
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

	userIdTemp := bm.Get(req.SessionId + "userId")
	userId := int(userIdTemp.([]uint8)[0])

	/* 处理数据 */
	// 查找订单信息
	o := orm.NewOrm()
	var userOrders []models.OrderHouse
	if req.Role == "landlord" {
		// 用户为房东, 找到自己目前已经发布了哪些房源
		var landlordHouses []models.House
		o.QueryTable("House").Filter("User__Id", userId).All(&landlordHouses)

		// 找到每一个房屋的id
		var houseIds []int
		for _, landlordHouse := range landlordHouses {
			houseIds = append(houseIds, landlordHouse.Id)
		}

		// 根觉房源id，找到对应订单
		o.QueryTable("OrderHouse").Filter("house__id__in", houseIds).OrderBy("Ctime").All(&userOrders)
	} else {
		// 用户为租客
		o.QueryTable("OrderHouse").Filter("User__Id", userId).OrderBy("Ctime").All(&userOrders)
	}

	var orderList []interface{}
	for _, userOrder := range userOrders {
		o.LoadRelated(&userOrder, "User")
		o.LoadRelated(&userOrder, "House")
		orderList = append(orderList, userOrder.ToOrderInfo())
	}

	// 更新 session
	bm.Put(req.SessionId+"userId", strconv.Itoa(userId), time.Second*3600)

	/* 返回数据 */
	rsp.Orders, _ = json.Marshal(orderList)

	return nil
}
