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
	"path"
	"time"

	postavatar "ihome/PostAvatar/proto/postavatar"
)

type Server struct{}

func (e *Server) PostAvatar(ctx context.Context, req *postavatar.Request, rsp *postavatar.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 获取文件的后缀名
	fileExt := path.Ext(req.FileName)

	// 连接 redis 数据库
	redisConfig, _ := json.Marshal(map[string]interface{}{
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
	// 获取 sessionId
	userIdTemp := bm.Get(req.SessionId + "userId")
	userId := int(userIdTemp.([]uint8)[0])

	// 从数据库中找到对应用户
	o := orm.NewOrm()
	var user models.User
	user.Id = userId

	/* 处理数据 */
	// 上传数据
	_, RemoteFileId, err := models.UploadByBuffer(req.Avatar, fileExt[1:])
	if err != nil {
		rsp.ErrNo = utils.RECODE_IOERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}
	user.AvatarUrl = RemoteFileId
	_, err = o.Update(&user, "AvatarUrl")
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 设置 session
	bm.Put(req.SessionId + "name", user.Name, time.Second * 3600)
	bm.Put(req.SessionId + "userId", string(user.Id), time.Second * 3600)
	bm.Put(req.SessionId + "mobile", user.Mobile, time.Second * 3600)

	/* 返回数据 */
	rsp.AvatarUrl = RemoteFileId
	return nil
}
