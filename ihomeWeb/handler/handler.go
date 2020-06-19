package handler

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	GETAREA "ihome/GetArea/proto/example"
	GETIMAGECD "ihome/GetImageCd/proto/example"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"image"
	"image/png"
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

// 获取图片验证码
func GetImageCd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 获取前端发送过来的图片唯一标识码
	uuid := ps.ByName("uuid")
	// 创建服务
	server := grpc.NewService()
	// 初始化服务
	server.Init()
	// 连接服务
	exampleCLient := GETIMAGECD.NewExampleService("go.micro.srv.GetImageCd", server.Client())
	rsp, err := exampleCLient.GetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid: uuid,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 处理图片信息
	var img image.RGBA
	img.Pix = []uint8(rsp.Pix)
	img.Stride = int(rsp.Stride)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)

	var image captcha.Image
	image.RGBA = &img

	// 将图片发送给前端
	png.Encode(w, image)
}
