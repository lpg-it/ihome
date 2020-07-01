package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"ihome/ihomeWeb/handler"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"time"

	postlogin "ihome/PostLogin/proto/postlogin"
)

type Example struct{}

func (e *Example) PostLogin(ctx context.Context, req *postlogin.Request, rsp *postlogin.Response) error {
	/* 初始化返回值 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 从数据库中查询是否有该用户
	o := orm.NewOrm()
	var user models.User
	err := o.QueryTable("User").Filter("Mobile", req.Mobile).One(&user)
	if err != nil {
		// 没有该用户
		rsp.ErrNo = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	/* 处理数据 */
	// 存在该用户，判断密码是否正确
	if user.PasswordHash != req.Password {
		// 密码不正确
		rsp.ErrNo = utils.RECODE_PWDERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	/* Redis 处理 */
	// 配置 redis 缓存信息
	redisConfig, _ := json.Marshal(map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	})
	// 连接 redis 数据库，创建对象
	bm, err := cache.NewCache("redis", string(redisConfig))
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 生成 sessionId
	sessionId := handler.GetMd5String(req.Mobile + req.Password)

	// 设置 session
	bm.Put(sessionId + "name", user.Name, time.Second * 3600)
	bm.Put(sessionId + "userId", string(user.Id), time.Second * 3600)
	bm.Put(sessionId + "mobile", user.Mobile, time.Second * 3600)

	/* 返回数据 */
	rsp.SessionId = sessionId
	return nil
}
