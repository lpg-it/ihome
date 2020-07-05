package handler

import (
	"context"
	"github.com/astaxie/beego/orm"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"path"
	"strconv"

	posthousesimage "ihome/PostHousesImage/proto/posthousesimage"
)

type Server struct{}

func (e *Server) PostHousesImage(ctx context.Context, req *posthousesimage.Request, rsp *posthousesimage.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */

	/* 处理数据 */
	// 将获取到的图片数据存入 fastDFS 中
	_, remoteFileId, err := models.UploadByBuffer(req.Image, path.Ext(req.FileName)[1:])
	if err != nil {
		rsp.ErrNo = utils.RECODE_IOERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 存入 数据库中
	o := orm.NewOrm()
	var house models.House
	house.Id, _ = strconv.Atoi(req.HouseId)
	_ = o.Read(&house)

	if house.IndexImageUrl == "" {
		// 主图为空则设置此图片为主图
		house.IndexImageUrl = remoteFileId
	}
	// 把此图片添加到 HouseImage 中
	var houseImage models.HouseImage
	houseImage.House = &house
	_ = o.Read(&houseImage)
	house.Images = append(house.Images, &houseImage)
	// 插入图片
	_, err = o.Insert(&houseImage)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	// 对 house 表进行更新
	_, err = o.Update(&house)
	if err != nil {
		rsp.ErrNo = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	/* 返回数据 */
	rsp.Url = remoteFileId
	return nil
}
