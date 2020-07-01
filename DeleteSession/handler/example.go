package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	"ihome/ihomeWeb/utils"

	deletesession "ihome/DeleteSession/proto/deletesession"
)

type Example struct{}

func (e *Example) DeleteSession(ctx context.Context, req *deletesession.Request, rsp *deletesession.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 获取 sessionId
	//req.SessionId

	/* 处理数据 */
	// 连接 redis 数据库
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

	// 删除缓存中的数据
	bm.Delete(req.SessionId + "name")
	bm.Delete(req.SessionId + "userId")
	bm.Delete(req.SessionId + "mobile")

	return nil
}
