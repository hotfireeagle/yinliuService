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

/**
** 查找一个appID所关联的banner列表数据
 */
func FindBannersByAppId(ctx *fiber.Ctx) error {
	appId := ctx.Params("appId") // 应用ID

	bannerIds := service.FindRelatedBannersByAppId(appId)
	var result []model.BannerJson

	if len(bannerIds) == 0 {
		res := model.IResponse{Code: model.Ok, Data: &result, Msg: "该应用并不存在banner"}
		return ctx.JSON(&res)
	}

	banners := service.FindBannersByIds(bannerIds)
	if len(*banners) == 0 {
		res := model.IResponse{Code: model.Ok, Data: &result, Msg: "指定bannerId在数据库并不存在"}
		return ctx.JSON(&res)
	}

	for _, bannerObj := range *banners {
		bjObj := model.BannerJson{
			Id:          bannerObj.Id,
			Src:         bannerObj.Src,
			RedirectUrl: bannerObj.RedirectUrl,
			Created:     bannerObj.Created,
		}
		result = append(result, bjObj)
	}

	return ctx.JSON(model.IResponse{Code: model.Ok, Data: &result})
}

/**
** 删除banner
 */
func DeleteBannerByBannerId(ctx *fiber.Ctx) error {
	bannerId := ctx.Params("bannerId")
	result := service.DeleteBannerService(bannerId)
	if result.Error != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "删除失败"}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Code: model.Ok}
	return ctx.JSON(&okRes)
}

func PatchBannerByBannerId(ctx *fiber.Ctx) error {
	bannerId := ctx.Params("bannerId")
	var bannerJson model.BannerJson
	err := ctx.BodyParser(&bannerJson)
	if err != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "JSON解析错误"}
		return ctx.JSON(&errRes)
	}
	bannerObj := model.Banner{
		Src:         bannerJson.Src,
		RedirectUrl: bannerJson.RedirectUrl,
	}
	result := service.PatchBannerService(bannerId, &bannerObj)
	if result.Error != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "更新失败"}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Code: model.Ok}
	return ctx.JSON(&okRes)
}
