package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
	"yinliuService/model"
	"yinliuService/service"
	"yinliuService/utils"
)

/**
** 后台用户进行登录
 */
func OpsUserDoLogin(ctx *fiber.Ctx) error {
	isPass, nextRes := emailAndUserVerifyHandler(ctx)

	if !isPass {
		return ctx.JSON(&nextRes)
	}

	user := nextRes.(model.User)

	hashedEmail := utils.Sha256(user.Phone)
	hashedPassword := utils.Sha256(user.Password)

	// 利用邮箱去查数据
	matchUser := service.FindUserByPhoneService(hashedEmail)
	if matchUser.Password != hashedPassword {
		errRes := model.IResponse{Code: model.Err, Msg: "密码错误"}
		return ctx.JSON(&errRes)
	}

	tokenCipher, err := generateToken(hashedEmail)

	if err != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "生成token失败"}
		return ctx.JSON(&errRes)
	}

	successRes := model.IResponse{Code: model.Ok, Msg: "", Data: tokenCipher}
	return ctx.JSON(&successRes)
}

/**
*** 创建用户
 */
func OpsUserDoCreate(ctx *fiber.Ctx) error {
	isPass, nextRes := emailAndUserVerifyHandler(ctx)
	if !isPass {
		return ctx.JSON(&nextRes)
	}

	user := nextRes.(model.User)

	hashedPhone := utils.Sha256(user.Phone)

	if service.IsPhoneExists("phone", hashedPhone) {
		errRes := model.IResponse{Code: model.Err, Msg: "该手机号已被注册"}
		return ctx.JSON(&errRes)
	}

	hashedPassword := utils.Sha256(user.Password)

	user.Password = hashedPassword
	user.Phone = hashedPhone

	result := service.CreateUserService(&user)
	if result.Error != nil {
		res := model.IResponse{Code: model.Err, Msg: "新建用户时发生错误"}
		return ctx.JSON(&res)
	}

	okRes := model.IResponse{Code: model.Ok, Msg: ""}
	return ctx.JSON(&okRes)
}

// 判断post参数中邮箱以及密码是否合法
func emailAndUserVerifyHandler(ctx *fiber.Ctx) (isPass bool, nextRes interface{}) {
	var user model.User
	err := ctx.BodyParser(&user)

	isPass = false // 默认false

	if err != nil {
		nextRes = model.IResponse{Code: model.Err, Msg: "请传入正确的JSON格式数据"}
		return
	}

	if user.Phone == "" {
		nextRes = model.IResponse{Code: model.Err, Msg: "手机号不能为空"}
		return
	}

	if user.Password == "" {
		nextRes = model.IResponse{Code: model.Err, Msg: "密码不能为空"}
		return
	}

	if !utils.VerifyMobilePhone(user.Phone) {
		nextRes = model.IResponse{Code: model.Err, Msg: "手机号格式错误"}
		return
	}

	isPass = true // 上面都不满足
	nextRes = user

	return
}

// 根据加密的邮箱号去生成token
func generateToken(encryptEmail string) (string, error) {
	userToken := model.UserToken{Uid: encryptEmail, Exp: time.Now().Add(time.Hour * 24)}
	jsonByteSlice, err := json.Marshal(userToken)
	tokenCipherByte := utils.AesEncode(jsonByteSlice)
	tokenCipher := utils.ByteSlice2Str(tokenCipherByte)
	return tokenCipher, err
}
