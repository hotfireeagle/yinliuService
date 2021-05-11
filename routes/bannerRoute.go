package routes

import (
	"github.com/gofiber/fiber/v2"
	"yinliuService/model"
	"yinliuService/service"
	"yinliuService/utils"
)

/**
** 创建新banner
 */
func CreateNewBannerRoute(ctx *fiber.Ctx) error {
	var createBannerJson model.CreateBannerJson
	err := ctx.BodyParser(&createBannerJson)

	if err != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "请传入正确的JSON格式数据"}
		return ctx.JSON(&errRes)
	}

	if createBannerJson.Src == "" {
		errRes := model.IResponse{Code: model.Err, Msg: "banner的链接必传"}
		return ctx.JSON(&errRes)
	}

	if len(createBannerJson.AppIds) == 0 {
		errRes := model.IResponse{Code: model.Err, Msg: "必须指定应用"}
		return ctx.JSON(&errRes)
	}

	apps := service.FindAppsByIds(createBannerJson.AppIds)

	if apps == nil || len(*apps) == 0 {
		errRes := model.IResponse{Code: model.Err, Msg: "所传入的appId均无效"}
		return ctx.JSON(&errRes)
	}

	id := utils.GenerateRandomString(16)
	bannerObj := model.Banner{
		Id:          id,
		Src:         createBannerJson.Src,
		RedirectUrl: createBannerJson.RedirectUrl,
		Apps:        *apps,
	}

	result := service.CreateBannerService(&bannerObj)
	if result.Error != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "存数据库时发生错误"}
		return ctx.JSON(&errRes)
	}

	okRes := model.IResponse{Code: model.Ok}
	return ctx.JSON(&okRes)
}
