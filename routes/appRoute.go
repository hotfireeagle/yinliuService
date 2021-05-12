package routes

import (
	"github.com/gofiber/fiber/v2"
	"yinliuService/model"
	"yinliuService/service"
	"yinliuService/utils"
)

/**
** 创建新app的路由方法
 */
func CreateNewAppRoute(ctx *fiber.Ctx) error {
	var app model.App
	err := ctx.BodyParser(&app)
	if err != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "请传入正确的json格式数据"}
		return ctx.JSON(&errRes)
	}
	if app.Name == "" {
		errRes := model.IResponse{Code: model.Err, Msg: "请指定应用名"}
		return ctx.JSON(&errRes)
	}
	id := utils.GenerateRandomString(16)
	app.Id = id

	result := service.CreateAppService(&app)
	if result.Error != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "保存到数据库发生错误"}
		return ctx.JSON(&errRes)
	}

	okRes := model.IResponse{Code: model.Ok}
	return ctx.JSON(&okRes)
}
