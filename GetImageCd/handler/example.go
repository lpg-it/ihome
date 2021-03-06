package handler

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"ihome/ihomeWeb/utils"
	"image/color"
	"time"

	getimagecd "ihome/GetImageCd/proto/getimagecd"
)

type Server struct{}

func (e *Server) GetImageCd(ctx context.Context, req *getimagecd.Request, rsp *getimagecd.Response) error {
	// 成功返回
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

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

	rsp.Stride = int64(c.Stride)
	rsp.Pix = []byte(c.Pix)
	rsp.Max = &example.Response_Point{X: int64(c.Rect.Max.X), Y: int64(c.Rect.Max.Y)}
	rsp.Min = &example.Response_Point{X: int64(c.Rect.Min.X), Y: int64(c.Rect.Min.Y)}

	// 存储到缓存中： 将 uuid 与 验证码的文本形式 存储在 redis 中
	// 初始化缓存全局变量的对象
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

	// 验证码进行 10 分钟的缓存
	_ = bm.Put(req.Uuid, str, 600*time.Second)

	return nil
}
