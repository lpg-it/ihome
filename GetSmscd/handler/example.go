package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"github.com/micro/go-log"
	example "ihome/GetSmscd/proto/example"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"math/rand"
	"time"
)

type Example struct{}

func (e *Example) GetSmscd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	// 初始化返回正确的返回值
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	// 验证手机号是否已经存在于数据库中
	o := orm.NewOrm()
	var user models.User
	user.Mobile = req.Mobile
	err := o.Read(&user)
	if err == nil {
		// 查到存在此手机号，返回 用户已存在
		rsp.ErrNo = utils.RECODE_DATAEXIST
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 验证 uuid 缓存
	// 连接 redis 数据库
	redisConf := map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	}
	redisConfJson, _ := json.Marshal(redisConf)
	// 连接 redis 数据库， 创建句柄
	bm, err := cache.NewCache("redis", string(redisConfJson))
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	uuidCache := bm.Get(req.Uuid)
	if uuidCache == nil {
		// 缓存中没有 uuid
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	uuidStr, _ := redis.String(uuidCache, nil)
	if req.Text != uuidStr {
		// 图片验证码错误
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// ============    设置 短信    ============
	// 生成一个随机数， 用于短信验证码
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomCode := r.Intn(99999) + 10001

	fmt.Println("短信验证码：", randomCode)

	// 通过手机号对验证短信进行缓存
	err = bm.Put(req.Mobile, randomCode, time.Second * 600)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Example) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
	log.Logf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&example.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
