package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-log"
	"ihome/ihomeWeb/handler"
	"net/http"

	"github.com/micro/go-web"
	_ "ihome/ihomeWeb/models"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.ihomeWeb"),
		web.Version("latest"),
		web.Address(":8080"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("html"))

	// register html handler
	service.Handle("/", rou)
	// 获取地区信息
	rou.GET("/api/v1.0/areas", handler.GetArea)
	// 获取图片验证码
	rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCd)
	// 获取短信验证码
	rou.GET("/api/v1.0/smscode/:mobile", handler.GetSmscd)
	// 注册
	rou.POST("/api/v1.0/users", handler.PostRet)
	// 获取 session
	rou.GET("/api/v1.0/session", handler.GetSession)
	// 登录
	rou.POST("/api/v1.0/sessions", handler.PostLogin)
	// 退出登录
	rou.DELETE("/api/v1.0/session", handler.DeleteSession)
	// 获取用户信息
	rou.GET("/api/v1.0/user", handler.GetUserInfo)
	// 上传用户头像
	rou.POST("/api/v1.0/user/avatar", handler.PostAvatar)
	// 更新用户名
	rou.PUT("/api/v1.0/user/name", handler.PutUserInfo)
	// 实名认证检查
	rou.GET("/api/v1.0/user/auth", handler.GetUserAuth)
	// 更新实名认证信息
	rou.POST("/api/v1.0/user/auth", handler.PostUserAuth)
	// 获取当前用户已发布房源信息
	rou.GET("/api/v1.0/user/houses", handler.GetUserHouses)
	// 发布房源信息
	rou.POST("/api/v1.0/houses", handler.PostHouses)
	// 上传房屋图片
	rou.POST("/api/v1.0/houses/:id/images", handler.PostHousesImage)
	// 获取房源详细信息
	rou.GET("/api/v1.0/houses/:id", handler.GetHouseInfo)
	// 获取首页轮播图
	rou.GET("/api/v1.0/house/index", handler.GetIndex)
	// 搜索房源
	rou.GET("/api/v1.0/houses", handler.GetHouses)
	// 发布订单
	rou.POST("/api/v1.0/orders", handler.PostOrders)
	// 查看 房东/租客 订单信息
	rou.GET("/api/v1.0/user/orders", handler.GetUserOrder)
	// 房东同意/拒绝订单
	rou.PUT("/api/v1.0/orders/:id/status", handler.PutOrders)
	// 用户评价订单
	rou.PUT("/api/v1.0/orders/:id/comment", handler.PutComment)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
