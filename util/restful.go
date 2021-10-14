package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"yinliuService/model"
	"yinliuService/tool"
)

// 检查客户端传过来的JSON数据是否是合法的JSON
func IsInvalidJson(ctx *fiber.Ctx, data interface{}) bool {
	if err := ctx.BodyParser(data); err != nil {
		tool.ErrLog(err.Error())
		res := model.IResponse{Status: model.Err, Msg: "无效JSON", ErrLog: err.Error()}
		ctx.JSON(&res)
		return true
	}
	return false
}

// 检查客户端传过来的数据是否字段校验合格
func IsInvalidData(ctx *fiber.Ctx, data interface{}) bool {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		tool.ErrLog(err.Error())
		res := model.IResponse{Status: model.Err, Msg: "参数错误", ErrLog: err.Error()}
		ctx.JSON(&res)
		return true
	}
	return false
}
