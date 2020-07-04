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

	postorders "ihome/PostOrders/proto/postorders"
)

type Server struct{}

func (e *Server) PostOrders(ctx context.Context, req *postorders.Request, rsp *postorders.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 获取 sessionId，以查找对应用户id
	// 连接 redis
	redisConfig, _ := json.Marshal(map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	})
	bm, err := cache.NewCache("redis", string(redisConfig))
	if err != nil {
		rsp.ErrNo = utils.RECODE_SESSIONERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	userIdTemp := bm.Get(req.SessionId + "userId")
	userId := int(userIdTemp.([]uint8)[0])

	// 获取订单信息
	orderInfo := make(map[string]interface{})
	err = json.Unmarshal(req.OrderInfo, &orderInfo)
	if err != nil {
		rsp.ErrNo = utils.RECODE_REQERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	/* 处理数据 */
	// 检验数据合法性
	// 简单的进行判断是否为空
	if orderInfo["house_id"] == "" || orderInfo["start_date"] == "" || orderInfo["end_date"] == "" {
		rsp.ErrNo = utils.RECODE_REQERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 结束日期要在开始日期之后
	// 格式化日期
	startDate, _ := time.Parse("2006-01-02 15:04:05", orderInfo["start_time"].(string)+"00:00:00")
	endDate, _ := time.Parse("2006-01-02 15:04:05", orderInfo["end_time"].(string)+"00:00:00")
	if endDate.Before(startDate) {
		rsp.ErrNo = utils.RECODE_ROLEERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 计算出入住的天数
	days := endDate.Sub(startDate).Hours()/24 + 1

	// 根据订单信息中的房屋id，找到关联的房屋信息
	var house models.House
	o := orm.NewOrm()
	house.Id, _ = strconv.Atoi(orderInfo["house_id"].(string))
	o.Read(&house)

	// 确保当前用户不是房源信息的发布者
	o.LoadRelated(&house, "User")
	if userId == house.User.Id {
		rsp.ErrNo = utils.RECODE_ROLEERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 添加订单信息
	/*Amount、Status、Comment
	Ctime、Credit*/
	var user models.User
	user.Id = userId
	o.Read(&user)

	var order models.OrderHouse
	order.House = &house
	order.User = &user
	order.BeginDate = startDate
	order.EndDate = endDate
	order.Days = int(days)
	order.HousePrice = house.Price
	order.Amount = int(days) * house.Price
	order.Status = models.OrderStatusWaitAccept
	order.Credit = false

	_, err = o.Insert(&order)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	bm.Put(req.SessionId+"userId", userId, time.Second*3600)

	/* 返回数据 */
	rsp.OrderId = strconv.Itoa(order.Id)

	return nil
}
