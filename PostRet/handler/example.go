package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	postret "ihome/PostRet/proto/postret"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"strconv"
	"time"
)

type Example struct{}

// md5 加密
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func (e *Example) PostRet(ctx context.Context, req *postret.Request, rsp *postret.Response) error {
	// 初始化错误码
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	// 获取数据
	// 从缓存中获取数据
	redisConf := map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	}
	redisConfJson, _ := json.Marshal(redisConf)

	// 连接 redis
	bm, err := cache.NewCache("redis", string(redisConfJson))
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	smsCodeCache := bm.Get(req.Mobile) // []uint8 类型
	if smsCodeCache == nil {
		// 没有缓存手机号数据
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 对 缓存数据 进行解码
	var smsCodeTemp interface{}
	json.Unmarshal(smsCodeCache.([]byte), &smsCodeTemp)

	// 类型转换 缓存中的手机号
	smsCode := int(smsCodeTemp.(float64))

	reqSmsCode, _ := strconv.Atoi(req.SmsCode)
	if reqSmsCode != smsCode {
		// 请求的短信验证码与缓存中的短信验证码不一致
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 往数据库中注册用户
	o := orm.NewOrm()
	var user models.User
	user.Name = req.Mobile
	user.Mobile = req.Mobile
	user.PasswordHash = getMd5String(req.Password)
	// 向数据库中插入数据
	_, err = o.Insert(&user)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 生成 sessionId 保证唯一性, 存入缓存中
	h := getMd5String(req.Mobile + req.Password)
	// 返回给客户端 session
	rsp.SessionId = h

	bm.Put(h+"name", string(user.Name), 3600*time.Second)
	bm.Put(h+"user_id", string(user.Id), 3600*time.Second)
	bm.Put(h+"mobile", string(user.Mobile), 3600*time.Second)

	return nil
}
