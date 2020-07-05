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

	putuserinfo "ihome/PutUserInfo/proto/putuserinfo"
)

type Server struct{}

func (e *Server) PutUserInfo(ctx context.Context, req *putuserinfo.Request, rsp *putuserinfo.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 获取 sessionId, 从 redis 中获取对应用户id， 在 mysql 中更新信息
	//sessionId := req.SessionId
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

	// 获取用户名
	//userName := req.UserName

	/* 处理数据 */
	// 去数据库更新用户名
	var user models.User
	o := orm.NewOrm()
	user.Id = userId
	user.Name = req.UserName
	_, err = o.Update(&user, "Name")
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 更新 session
	_ = bm.Put(req.SessionId+"userId", user.Id, time.Second*3600)
	_ = bm.Put(req.SessionId+"name", user.Name, time.Second*3600)
	_ = bm.Put(req.SessionId+"mobile", user.Mobile, time.Second*3600)

	/* 返回数据 */
	rsp.UserName = user.Name
	return nil
}
