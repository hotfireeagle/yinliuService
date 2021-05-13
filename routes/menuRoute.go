package routes

import (
	"github.com/gofiber/fiber/v2"
	"yinliuService/model"
	"yinliuService/service"
	"yinliuService/utils"
)

/**
** 新增菜单栏的操作方法
 */
func CreateNewMenuRoute(ctx *fiber.Ctx) error {
	var jsonBody model.CreateNewMenuJson
	err := ctx.BodyParser(&jsonBody)
	if err != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "json解析错误"}
		return ctx.JSON(&errRes)
	}

	if len(jsonBody.AppIds) == 0 {
		errRes := model.IResponse{Code: model.Err, Msg: "appIds不能为空"}
		return ctx.JSON(&errRes)
	}

	if jsonBody.Text == "" {
		errRes := model.IResponse{Code: model.Err, Msg: "menu文字不能为空"}
		return ctx.JSON(&errRes)
	}

	if jsonBody.Icon == "" {
		errRes := model.IResponse{Code: model.Err, Msg: "menu图标不能为空"}
		return ctx.JSON(&errRes)
	}

	apps := service.FindAppsByIds(jsonBody.AppIds)
	if len(*apps) == 0 {
		errRes := model.IResponse{Code: model.Err, Msg: "appIds均无效"}
		return ctx.JSON(&errRes)
	}

	id := utils.GenerateRandomString(16)
	menuObj := model.Menu{
		Id:          id,
		Icon:        jsonBody.Icon,
		Text:        jsonBody.Text,
		RedirectUrl: jsonBody.RedirectUrl,
		Apps:        *apps,
	}
	result := service.CreateNewMenuService(&menuObj)
	if result.Error != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "存Menu时发生错误"}
		return ctx.JSON(&errRes)
	}

	okRes := model.IResponse{Code: model.Ok}
	return ctx.JSON(&okRes)
}

/**
** 根据应用ID找出其对应的所有菜单
 */
func FindMenusByAppId(ctx *fiber.Ctx) error {
	appId := ctx.Params("appId")

	menuIds := service.FindRelatedMenusByAppId(appId)
	result := make([]model.MenuJson, 0)

	if len(menuIds) == 0 {
		res := model.IResponse{Code: model.Ok, Data: &result, Msg: "该应用并不存在菜单数据"}
		return ctx.JSON(&res)
	}

	menus := service.FindMenusByIds(menuIds)
	if len(*menus) == 0 {
		res := model.IResponse{Code: model.Ok, Data: &result, Msg: "指定的menuIds在数据库并不存在"}
		return ctx.JSON(&res)
	}

	for _, menuObj := range *menus {
		bjObj := model.MenuJson{
			Id:          menuObj.Id,
			Icon:        menuObj.Icon,
			Text:        menuObj.Text,
			RedirectUrl: menuObj.RedirectUrl,
			Created:     menuObj.Created,
		}
		result = append(result, bjObj)
	}

	return ctx.JSON(model.IResponse{Code: model.Ok, Data: &result})
}

/**
** 删除指定ID的菜单
 */
func DeleteMenuByMenuId(ctx *fiber.Ctx) error {
	menuId := ctx.Params("menuId")
	result := service.DeleteMenuService(menuId)
	if result.Error != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "删除失败"}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Code: model.Ok}
	return ctx.JSON(&okRes)
}

/**
** 更新menu
 */
func PatchMenuByMenuId(ctx *fiber.Ctx) error {
	menuId := ctx.Params("menuId")
	var menuJson model.MenuJson
	err := ctx.BodyParser(&menuJson)
	if err != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "JSON解析错误"}
		return ctx.JSON(&errRes)
	}
	menuObj := model.Menu{
		Icon:        menuJson.Icon,
		RedirectUrl: menuJson.RedirectUrl,
		Text:        menuJson.Text,
	}
	result := service.PatchMenuService(menuId, &menuObj)
	if result.Error != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "更新失败"}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Code: model.Ok}
	return ctx.JSON(&okRes)
}
