package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	deletesession "ihome/DeleteSession/proto/deletesession"
	getarea "ihome/GetArea/proto/getarea"
	GETIMAGECD "ihome/GetImageCd/proto/example"
	getsession "ihome/GetSession/proto/getsession"
	GETSMSCD "ihome/GetSmscd/proto/example"
	getuserauth "ihome/GetUserAuth/proto/getuserauth"
	getuserinfo "ihome/GetUserInfo/proto/getuserinfo"
	postavatar "ihome/PostAvatar/proto/postavatar"
	postlogin "ihome/PostLogin/proto/postlogin"
	postret "ihome/PostRet/proto/postret"
	postuserauth "ihome/PostUserAuth/proto/postuserauth"
	putuserinfo "ihome/PutUserInfo/proto/putuserinfo"
	"ihome/ihomeWeb/models"
	"ihome/ihomeWeb/utils"
	"image"
	"image/png"
	"net/http"
	"regexp"
)

// md5 加密
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 获取地区信息
func GetArea(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 创建新的 gRPC 返回句柄
	server := grpc.NewService()
	// 服务初始化
	server.Init()

	// 创建获取地区的服务并且返回句柄
	exampleClient := getarea.NewExampleService("go.micro.srv.GetArea", server.Client())
	// 调用函数并且返回数据
	rsp, err := exampleClient.GetArea(context.TODO(), &getarea.Request{})
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
	// 获取数据
	// 获取 cookie
	userLogin, err := r.Cookie("userLogin")
	if err != nil {
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	// 存在就发送数据给服务
	// 创建服务
	service := grpc.NewService()
	service.Init()

	getSessionClient := getsession.NewExampleService("go.micro.srv.GetSession", service.Client())
	rsp, err := getSessionClient.GetSession(context.TODO(), &getsession.Request{
		SessionId: userLogin.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	// 将获取到的用户名给前端
	data := make(map[string]string)
	data["name"] = rsp.Data
	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
		"data":   data,
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
		"errNo":  utils.RECODE_OK,
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

// 获取手机短信验证码
func GetSmscd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// ============    获取数据    ============
	// 获取手机号
	mobile := ps.ByName("mobile")
	// 获取 uuid
	uuid := r.URL.Query()["id"][0]
	// 获取 图片验证码
	imageCode := r.URL.Query()["text"][0]

	// ============    处理数据    ============
	// 手机号 进行正则匹配
	// 创建正则对象
	regex := regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)
	isMobile := regex.MatchString(mobile)
	if isMobile == false {
		// 手机号格式错误, 返回
		response := map[string]interface{}{
			"errno":  utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}
		// 设置返回数据格式
		w.Header().Set("Content-Type", "application/json")
		// 将错误发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	// ============    创建服务    ============
	server := grpc.NewService()
	server.Init()

	// 调用服务
	exampleClient := GETSMSCD.NewExampleService("go.micro.srv.GetSmscd", server.Client())
	rsp, err := exampleClient.GetSmscd(context.TODO(), &GETSMSCD.Request{
		Mobile: mobile,
		Uuid:   uuid,
		Text:   imageCode,
	})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	// 返回
	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
	}
	w.Header().Set("Content-Type", "application/json")
	// 将数据返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
}

// 注册
func PostRet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 获取数据
	// 获取前端发送过来的 json 数据
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//for key, value := range request {
	//	fmt.Println("key: ", key)
	//	fmt.Println("value: ", value)
	//}

	// 验证数据
	// 判断是否为空
	if request["mobile"] == "" || request["password"] == "" || request["sms_code"] == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}
		// 如果不存在直接给前端返回
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	// 创建服务
	service := grpc.NewService()
	service.Init()

	postRetClient := postret.NewExampleService("go.micro.srv.PostRet", service.Client())
	rsp, err := postRetClient.PostRet(context.TODO(), &postret.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
		SmsCode:  request["sms_code"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
	}

	// 读取 cookie, 如果 cookie 不存在则创建
	cookie, err := r.Cookie("userLogin")
	if err != nil || cookie.Value == "" {
		// 创建 cookie
		cookie := http.Cookie{
			Name:   "userLogin",
			Value:  rsp.SessionId,
			Path:   "/",
			MaxAge: 600,
		}
		http.SetCookie(w, &cookie)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	return
}

// 登录
func PostLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/* 获取前端发送的数据 */
	var request map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	/* 处理数据 */
	// 判断账号密码是否为空
	if request["mobile"] == "" || request["password"] == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}
	/* 连接服务 */
	service := grpc.NewService()
	service.Init()

	postLoginClient := postlogin.NewPostLoginService("go.micro.srv.PostLogin", service.Client())
	rsp, err := postLoginClient.PostLogin(context.TODO(), &postlogin.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	// 获取cookie
	userLoginCookie, err := r.Cookie("userLogin")
	if err != nil || userLoginCookie.Value == "" {
		// 没有cookie，设置cookie
		userLoginCookie := http.Cookie{
			Name:   "userLogin",
			Value:  rsp.SessionId,
			Path:   "/",
			MaxAge: 600,
		}
		http.SetCookie(w, &userLoginCookie)
	}

	/* 返回数据 */
	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	return
}

// 退出登录
func DeleteSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/* 获取数据 */
	// 获取 sessionId
	userLoginSession, err := r.Cookie("userLogin")
	if err != nil {
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
	}

	// 连接 退出登录 服务
	service := grpc.NewService()
	service.Init()
	deleteSessionClient := deletesession.NewDeleteSessionService("go.micro.srv.DeleteSession", service.Client())
	rsp, err := deleteSessionClient.DeleteSession(context.TODO(), &deletesession.Request{
		SessionId: userLoginSession.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	/* 处理数据 */

	http.SetCookie(w, &http.Cookie{
		Name:   "userLogin",
		Path:   "/",
		MaxAge: -1,
	})

	/* 返回数据 */
	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	return
}

// 获取用户信息
func GetUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/* 获取数据 */
	// 获取 sessionId
	userLoginSession, err := r.Cookie("userLogin")
	if err != nil {
		// 获取 session 失败，直接返回
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	/* 处理数据 */
	// 连接 服务
	service := grpc.NewService()
	service.Init()

	getUserInfoClient := getuserinfo.NewGetUserInfoService("go.micro.srv.GetUserInfo", service.Client())
	rsp, err := getUserInfoClient.GetUserInfo(context.TODO(), &getuserinfo.Request{
		SessionId: userLoginSession.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	data := make(map[string]interface{})
	// 将从服务端得到的数据发送给前端
	data["user_id"] = rsp.UserId
	data["name"] = rsp.UserName
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

	/* 返回数据 */
	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	return
}

// 上传用户头像
func PostAvatar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/* 获取数据 */
	// 获取用户 sessionId, 查看登录信息
	userLoginSession, err := r.Cookie("userLogin")
	if err != nil {
		// 获取 session 失败, 返回前端数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	// 获取二进制图片，名字，大小
	avatarFile, avatarHeader, err := r.FormFile("avatar")
	if err != nil {
		response := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	/* 处理数据 */
	// 存储文件（二进制）
	fileBuffer := make([]byte, avatarHeader.Size)
	// 将文件读到 fileBuffer 里
	_, err = avatarFile.Read(fileBuffer)
	if err != nil {
		response := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	// 连接 上传头像 服务， 传入数据
	service := grpc.NewService()
	service.Init()

	postAvatarClient := postavatar.NewPostAvatarService("go.micro.srv.PostAvatar", service.Client())
	rsp, err := postAvatarClient.PostAvatar(context.TODO(), &postavatar.Request{
		SessionId: userLoginSession.Value,
		Avatar:    fileBuffer,
		FileName:  avatarHeader.Filename,
		FileSize:  avatarHeader.Size,
	})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	/* 返回数据 */
	// 给前端传输数据
	data := make(map[string]interface{})
	// url 拼接图片地址
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	return
}

// 更新用户名
func PutUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/* 获取数据 */
	// 获取前端提交的数据
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// 获取 sessionId
	userLoginSession, err := r.Cookie("userLogin")
	if err != nil {
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	/* 处理数据 */
	// 连接 更新用户名 服务
	service := grpc.NewService()
	service.Init()

	putUserInfoClient := putuserinfo.NewPutUserInfoService("go.micro.srv.PutUserInfo", service.Client())
	rsp, err := putUserInfoClient.PutUserInfo(context.TODO(), &putuserinfo.Request{
		SessionId: userLoginSession.Value,
		UserName:  request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	/* 返回数据 */
	// 接收回发的数据
	data := map[string]interface{}{
		"name": rsp.UserName,
	}

	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	return
}

// 获取用户实名认证
func GetUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/* 获取数据 */
	// 获取 sessionId
	userLoginSession, err := r.Cookie("userLogin")
	if err != nil {
		// 获取 session 失败，直接返回
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
		return
	}

	/* 处理数据 */
	// 连接 服务
	service := grpc.NewService()
	service.Init()

	getUserAuthClient := getuserauth.NewGetUserAuthService("go.micro.srv.GetUserAuth", service.Client())
	rsp, err := getUserAuthClient.GetUserAuth(context.TODO(), &getuserauth.Request{
		SessionId: userLoginSession.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	data := make(map[string]interface{})
	// 将从服务端得到的数据发送给前端
	data["user_id"] = rsp.UserId
	data["name"] = rsp.UserName
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

	/* 返回数据 */
	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	return
}

// 更新实名认证
func PostUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/* 获取数据 */
	// 获取前端提交的数据
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 获取 session
	userLoginSession, err := r.Cookie("userLogin")
	if err != nil {
		// 用户未登录
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&response); err != nil {
			http.Error(w, err.Error(), 503)
			return
		}
	}

	/* 处理数据 */
	// 连接 更新用户实名认证 服务
	service := grpc.NewService()
	service.Init()

	postUserAuthClient := postuserauth.NewPostUserAuthService("go.micro.srv.PostUserAuth", service.Client())
	rsp, err := postUserAuthClient.PostUserAuth(context.TODO(), &postuserauth.Request{
		SessionId: userLoginSession.Value,
		RealName:  request["real_name"].(string),
		IdCard:    request["id_card"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	/* 返回数据 */
	response := map[string]interface{}{
		"errno":  rsp.ErrNo,
		"errmsg": rsp.ErrMsg,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	return
}
