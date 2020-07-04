package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"strconv"

	gethouses "ihome/GetHouses/proto/gethouses"
)

type Server struct{}

func (e *Server) GetHouses(ctx context.Context, req *gethouses.Request, rsp *gethouses.Response) error {
	/* 初始化返回数据 */
	rsp.ErrNo = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)

	/* 获取数据 */
	// 地区id
	areaId := req.AreaId
	// 起始日期
	//startDate := req.StartDate
	// 结束日期
	//endDate := req.EndDate
	// 页码
	pageNumber, _ := strconv.Atoi(req.PageNumber)

	// 根据查询条件查询数据
	var houses []models.House
	o := orm.NewOrm()
	houseNum, err := o.QueryTable("House").Filter("Area__Id", areaId).All(&houses)
	if err != nil {
		rsp.ErrNo = utils.RECODE_PARAMERR
		rsp.ErrMsg = utils.RecodeText(rsp.ErrNo)
		return nil
	}

	/* 处理数据 */
	totalPage := (int(houseNum) / models.HouseListPageCapacity) + 1

	var houseList []interface{}
	for _, house := range houses {
		o.LoadRelated(&house, "Area")
		o.LoadRelated(&house, "User")
		o.LoadRelated(&house, "Images")
		o.LoadRelated(&house, "Facilities")

		houseList = append(houseList, house.ToHouseInfo())
	}

	/* 返回数据 */
	rsp.Houses, _ = json.Marshal(houseList)
	rsp.CurrentPage = int64(pageNumber)
	rsp.TotalPage = int64(totalPage)

	return nil
}
