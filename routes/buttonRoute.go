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
func CreateNewButtonRoute(ctx *fiber.Ctx) error {
	var jsonBody model.CreateNewButtonJson
	err := ctx.BodyParser(&jsonBody)
	if err != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "json解析错误"}
		return ctx.JSON(&errRes)
	}

	if len(jsonBody.AppIds) == 0 {
		errRes := model.IResponse{Code: model.Err, Msg: "appIds不能为空"}
		return ctx.JSON(&errRes)
	}

	if jsonBody.Icon == "" {
		errRes := model.IResponse{Code: model.Err, Msg: "Icon不能为空"}
		return ctx.JSON(&errRes)
	}

	if jsonBody.Title == "" {
		errRes := model.IResponse{Code: model.Err, Msg: "主标题不能为空"}
		return ctx.JSON(&errRes)
	}

	if jsonBody.Desc == "" {
		errRes := model.IResponse{Code: model.Err, Msg: "副标题不能为空"}
		return ctx.JSON(&errRes)
	}

	if jsonBody.BtnTxt == "" {
		errRes := model.IResponse{Code: model.Err, Msg: "按钮文案不能为空"}
		return ctx.JSON(&errRes)
	}

	apps := service.FindAppsByIds(jsonBody.AppIds)
	if len(*apps) == 0 {
		errRes := model.IResponse{Code: model.Err, Msg: "appIds均无效"}
		return ctx.JSON(&errRes)
	}

	id := utils.GenerateRandomString(16)
	buttonObj := model.Button{
		Id:          id,
		Icon:        jsonBody.Icon,
		Title:       jsonBody.Title,
		Desc:        jsonBody.Desc,
		BtnTxt:      jsonBody.BtnTxt,
		RedirectUrl: jsonBody.RedirectUrl,
		Apps:        *apps,
	}
	result := service.CreateNewButtonService(&buttonObj)
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
func FindButtonsByAppId(ctx *fiber.Ctx) error {
	appId := ctx.Params("appId")

	buttonIds := service.FindRelatedButtonsByAppId(appId)
	result := make([]model.ButtonJson, 0)

	if len(buttonIds) == 0 {
		res := model.IResponse{Code: model.Ok, Data: &result, Msg: "该应用并不存在按钮数据"}
		return ctx.JSON(&res)
	}

	buttons := service.FindButtonsByIds(buttonIds)
	if len(*buttons) == 0 {
		res := model.IResponse{Code: model.Ok, Data: &result, Msg: "指定的buttonIds在数据库并不存在"}
		return ctx.JSON(&res)
	}

	for _, buttonObj := range *buttons {
		bjObj := model.ButtonJson{
			Id:          buttonObj.Id,
			Icon:        buttonObj.Icon,
			Title:       buttonObj.Title,
			Desc:        buttonObj.Desc,
			BtnTxt:      buttonObj.BtnTxt,
			RedirectUrl: buttonObj.RedirectUrl,
			Created:     buttonObj.Created,
		}
		result = append(result, bjObj)
	}

	return ctx.JSON(model.IResponse{Code: model.Ok, Data: &result})
}

/**
** 删除指定ID的菜单
 */
func DeleteButtonByButtonId(ctx *fiber.Ctx) error {
	buttonId := ctx.Params("buttonId")
	result := service.DeleteButtonService(buttonId)
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
func PatchButtonByButtonId(ctx *fiber.Ctx) error {
	buttonId := ctx.Params("buttonId")
	var buttonJson model.ButtonJson
	err := ctx.BodyParser(&buttonJson)
	if err != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "JSON解析错误"}
		return ctx.JSON(&errRes)
	}
	buttonObj := model.Button{
		Icon:        buttonJson.Icon,
		Title:       buttonJson.Title,
		Desc:        buttonJson.Desc,
		BtnTxt:      buttonJson.BtnTxt,
		RedirectUrl: buttonJson.RedirectUrl,
	}
	result := service.PatchButtonService(buttonId, &buttonObj)
	if result.Error != nil {
		errRes := model.IResponse{Code: model.Err, Msg: "更新失败"}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Code: model.Ok}
	return ctx.JSON(&okRes)
}
