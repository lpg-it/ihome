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

	putorders "ihome/PutOrders/proto/putorders"
)

type Server struct{}

func (e *Server) PutOrders(ctx context.Context, req *putorders.Request, rsp *putorders.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 获取 sessionId， 找到对应用户id
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
	userId := bm.Get(req.SessionId + "userId")

	// 获取订单id
	orderId, _ := strconv.Atoi(req.OrderId)

	// 获取 action
	action := req.Action

	/* 处理数据 */
	// 查找订单表， 找到该订单， 并确定当前状态
	o := orm.NewOrm()
	var userOrder models.OrderHouse
	userOrder.Id = orderId
	err = o.QueryTable("OrderHouse").Filter("Id", orderId).Filter("Status", models.OrderStatusWaitAccept).One(&userOrder)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	o.LoadRelated(&userOrder, "House")
	// 检验订单的 userid 是否是当前用户的 id
	if userId != userOrder.House.User.Id {
		rsp.ErrNo = utils.RECODE_DATAERR
		rsp.ErrMsg = "订单用户不匹配"
		return nil

	}

	// 修改 action
	if action == "accept" {
		// 如果是接受订单， 修改订单状态为 带评价状态
		userOrder.Status = models.OrderStatusWaitComment
	} else if action == "reject" {
		// 如果是拒绝订单， 修改订单状态为 拒绝状态， 并获得拒绝原因
		userOrder.Status = models.OrderStatusRejected
		userOrder.Comment = req.Action
	}

	// 更新数据
	if _, err := o.Update(&userOrder); err != nil {
		rsp.ErrNo = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 设置 session
	_ = bm.Put(req.SessionId+"userId", userId.(string), time.Second*3600)

	return nil
}
