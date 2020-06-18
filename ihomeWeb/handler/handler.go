package handler

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	GETAREA "ihome/GetArea/proto/example"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"net/http"
)

// 获取地区信息
func GetArea(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 创建新的 gRPC 返回句柄
	server := grpc.NewService()
	// 服务初始化
	server.Init()

	// 创建获取地区的服务并且返回句柄
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea", server.Client())
	// 调用函数并且返回数据
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		return
	}
	// 创建返回类型的切片
	var areas []models.Area
	// 循环读取服务返回的数据
	for _, value := range rsp.Data {
		areas = append(areas, models.Area{Id: int(value.Aid), Name: value.Aname})
	}
	// 创建返回数据 map
	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
		"data":   areas,
	}
	// 注意
	w.Header().Set("Content-Type", "application/json")

	// 将返回的数据 map 发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
}

// 获取 session
func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 创建返回的数据
	response := map[string]interface{}{
		"errNo": utils.RECODE_SESSIONERR,
		"errMsg": utils.RecodeText(utils.RECODE_SESSIONERR),
	}
	w.Header().Set("Content-Type", "application/json")
	// 将返回的数据 map 发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
}

// 获取首页轮播图
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 创建返回的数据
	response := map[string]interface{}{
		"errNo": utils.RECODE_OK,
		"errMsg": utils.RecodeText(utils.RECODE_OK),
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
}