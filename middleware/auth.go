package middleware

import (
	"encoding/hex"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
	"yinliuService/model"
	"yinliuService/service"
	"yinliuService/tool"
)

/**
** token认证中间件
 */

func TokenAuth(ctx *fiber.Ctx) error {
	token := model.UserToken{}
	tokenCipherSlice := ctx.Request().Header.Peek("token")
	tokenCipherStr := string(tokenCipherSlice)

	if tokenCipherStr == "" {
		errRes := model.IResponse{Status: model.UnLogin, Msg: "请先登录"}
		return ctx.JSON(&errRes)
	}

	hexTokenCipherSlice, err := hex.DecodeString(tokenCipherStr)
	if err != nil {
		errRes := model.IResponse{Status: model.UnLogin, Msg: "token解码失败"}
		return ctx.JSON(&errRes)
	}

	tokenOriginSlice := tool.AesDecode(hexTokenCipherSlice)
	err = json.Unmarshal(tokenOriginSlice, &token)
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "JSON解析错误"}
		return ctx.JSON(&errRes)
	}

	if !service.FindUserByIdService(token.Uid) {
		errRes := model.IResponse{Status: model.UnLogin, Msg: "用户不存在"}
		return ctx.JSON(&errRes)
	}

	if token.Exp.Before(time.Now()) {
		errRes := model.IResponse{Status: model.LoginOverdue, Msg: "登录已过期，请重新登录"}
		return ctx.JSON(&errRes)
	}

	ctx.Request().Header.Set("uid", token.Uid)
	return ctx.Next()
}
