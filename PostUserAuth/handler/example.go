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
	"time"

	postuserauth "ihome/PostUserAuth/proto/postuserauth"
)

type Server struct{}

func (e *Server) PostUserAuth(ctx context.Context, req *postuserauth.Request, rsp *postuserauth.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 获取 sessionId: req.SessionId
	// 获取 真实姓名: req.RealName
	// 获取 身份证号: req.IdCard
	// 连接 redis，获取到对应的用户
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

	// 获取 用户 id
	userIdTemp := bm.Get(req.SessionId + "userId")
	userId := int(userIdTemp.([]uint8)[0])

	/* 处理数据 */
	// 去数据库更新信息
	var user models.User
	o := orm.NewOrm()
	user.Id = userId
	user.RealName = req.RealName
	user.IdCard = req.IdCard
	_, err = o.Update(&user, "RealName", "IdCard")
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	/* 返回数据 */
	bm.Put(req.SessionId+"userId", user.Id, time.Second*3600)
	bm.Put(req.SessionId+"name", user.Name, time.Second*3600)
	bm.Put(req.SessionId+"mobile", user.Mobile, time.Second*3600)

	return nil
}
