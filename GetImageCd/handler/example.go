package handler

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego/cache"
	"ihome/ihomeWeb/utils"
	"image/color"
	"time"

	"github.com/micro/go-log"

	example "ihome/GetImageCd/proto/example"
)

type Example struct{}

func (e *Example) GetImageCd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	// 创建一个对象
	cap := captcha.New()
	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}
	// 设置图片的大小
	cap.SetSize(90, 41)
	// 设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 设置前景色，  可以设置多个 随机替换文字颜色， 默认黑色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色， 可以设置多个 随机替换背景色， 默认 白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	// 生成图片， 返回图片和字符串(图片内容的文本形式)
	img, str := cap.Create(4, captcha.NUM)
	// 解引用
	b := *img
	c := *(b.RGBA)

	// 成功返回
	rsp.ErrNO = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNO)

	rsp.Stride = int64(c.Stride)
	rsp.Pix = []byte(c.Pix)
	rsp.Max = &example.Response_Point{X: int64(c.Rect.Max.X), Y: int64(c.Rect.Max.Y)}
	rsp.Min = &example.Response_Point{X: int64(c.Rect.Min.X), Y: int64(c.Rect.Min.Y)}

	// 存储到缓存中： 将 uuid 与 验证码的文本形式 存储在 redis 中
	// 初始化缓存全局变量的对象
	redisConf := map[string]interface{}{
		"key":   "ihome",
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNUm": utils.G_redis_dbnum,
	}
	redisConfJson, _ := json.Marshal(redisConf)

	// 连接 redis 创建对象
	bm, err := cache.NewCache("redis", string(redisConfJson))
	if err != nil {
		rsp.ErrNO = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNO)
		return nil
	}
	// 验证码进行 10 分钟的缓存
	bm.Put(req.Uuid, str, 600*time.Second)

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
