package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"strconv"

	putcomment "ihome/PutComment/proto/putcomment"
)

type Server struct{}

func (e *Server) PutComment(ctx context.Context, req *putcomment.Request, rsp *putcomment.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 获取 用户id
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
	userId := int(bm.Get(req.SessionId + "userId").([]uint8)[0])

	// 获取 订单id
	orderId, _ := strconv.Atoi(req.OrderId)

	// 获取 订单评论
	orderComment := req.OrderComment

	/* 处理数据 */
	// 检验评论是否合法：判空
	if orderComment == "" {
		rsp.ErrNo = utils.RECODE_PARAMERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 根据 订单id 找到该订单
	o := orm.NewOrm()
	var order models.OrderHouse
	// 状态必须为待评价
	if err := o.QueryTable("OrderHouse").Filter("Id", orderId).Filter("Status", models.OrderStatusWaitComment).One(&order); err != nil {
		rsp.ErrNo = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 查询订单关联的用户信息， 确保与登录用户是同一人
	if _, err := o.LoadRelated(&order, "User"); err != nil {
		rsp.ErrNo = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	if order.User.Id != userId {
		rsp.ErrNo = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 关联查询 房源信息
	if _, err := o.LoadRelated(&order, "House"); err != nil {
		rsp.ErrNo = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 增加订单评论， 并且更新 状态为 Complete
	order.Comment = orderComment
	order.Status = models.OrderStatusComplete

	// 房屋的成交量 +1
	order.House.OrderCount++

	// 更新内容到数据库
	if _, err := o.Update(&order); err != nil {
		rsp.ErrNo = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 更新 房源缓存（已修改房源的订单数量）
	_ = bm.Delete("houseInfo" + strconv.Itoa(order.House.Id))

	/* 返回数据 */

	return nil
}
