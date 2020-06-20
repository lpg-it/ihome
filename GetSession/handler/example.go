package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	getsession "ihome/GetSession/proto/getsession"
	"ihome/ihomeWeb/utils"
)

type Example struct{}

func (e *Example) GetSession(ctx context.Context, req *getsession.Request, rsp *getsession.Response) error {
	// 初始化
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	// 获取前端的 cookie
	redisConf := map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	}
	redisConfJson, _ := json.Marshal(redisConf)
	bm, err := cache.NewCache("redis", string(redisConfJson))
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	userNameSession := bm.Get(req.SessionId + "name")
	userName, err := redis.String(userNameSession, nil)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 获取到了 session
	rsp.Data = userName

	return nil
}
