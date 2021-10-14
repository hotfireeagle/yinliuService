package route

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
	"yinliuService/model"
	"yinliuService/service"
	"yinliuService/tool"
	"yinliuService/util"
)

// 后台用户进行登录操作
func OpsUserDoLogin(ctx *fiber.Ctx) error {
	user := new(model.User)

	// 先对传参进行校验
	if util.IsInvalidJson(ctx, user) {
		return nil
	}
	if util.IsInvalidData(ctx, user) {
		return nil
	}

	// 对手机以及密码进行hash
	hashedPassword := tool.Sha256(user.Password)

	// 利用手机号进行数据查询
	matchUser, err := service.FindUserByPhoneService(user.Phone)

	if err != nil {
		tool.ErrLog(err)
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}

	if matchUser.ID == "" {
		errRes := model.IResponse{Status: model.Err, Msg: "该用户未注册"}
		return ctx.JSON(&errRes)
	}

	if matchUser.Password != hashedPassword {
		errRes := model.IResponse{Status: model.Err, Msg: "密码错误"}
		return ctx.JSON(&errRes)
	}

	tokenCipher, err := generateToken(matchUser.ID) // 根据ID生成token

	if err != nil {
		tool.ErrLog(err)
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}

	successRes := model.IResponse{Status: model.Success, Data: tokenCipher}
	return ctx.JSON(&successRes)
}

// 创建新用户
func OpsUserDoCreate(ctx *fiber.Ctx) error {
	user := new(model.User)
	if util.IsInvalidJson(ctx, user) {
		return nil
	}
	if util.IsInvalidData(ctx, user) {
		return nil
	}


	isPhoneUsed := service.FindUserIsExistsByPhoneService(user.Phone)
	if isPhoneUsed {
		errRes := model.IResponse{Status: model.Err, Msg: "该手机号已被注册"}
		return ctx.JSON(&errRes)
	}

	hashedPassword := tool.Sha256(user.Password)

	user.Password = hashedPassword
	user.CreatedAt = time.Now()

	user.GenerateUUID()
	err := service.CreateUserService(user)
	if err != nil {
		tool.ErrLog(err)
		res := model.IResponse{Status: model.Err, Msg: "系统错误", ErrLog: err.Error()}
		return ctx.JSON(&res)
	}

	okRes := model.IResponse{Status: model.Success}
	return ctx.JSON(&okRes)
}


// 根据uid生成token
func generateToken(encryptEmail string) (string, error) {
	userToken := model.UserToken{Uid: encryptEmail, Exp: time.Now().Add(time.Hour * 24)}
	jsonByteSlice, err := json.Marshal(userToken)
	tokenCipherByte := tool.AesEncode(jsonByteSlice)
	tokenCipher := tool.ByteSlice2Str(tokenCipherByte)
	return tokenCipher, err
}
