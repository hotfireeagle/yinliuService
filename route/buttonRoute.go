package route

import (
	"github.com/gofiber/fiber/v2"
	"yinliuService/model"
	"yinliuService/service"
	"yinliuService/util"
)

// 新增按钮
func createNewButtonRoute(ctx *fiber.Ctx) error {
	button := new(model.Button)
	if util.IsInvalidJson(ctx, button) {
		return nil
	}
	if util.IsInvalidData(ctx, button) {
		return nil
	}
	err := service.CreateNewButtonService(button)
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success}
	return ctx.JSON(&okRes)
}

// 找出所有按钮
func findButtonsByAppId(ctx *fiber.Ctx) error {
	results, err := service.FindButtons()
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success, Data: results}
	return ctx.JSON(&okRes)
}

// 删除指定的按钮
func deleteButtonByButtonId(ctx *fiber.Ctx) error {
	err := service.DeleteButtonService(ctx.Params("id"))
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success}
	return ctx.JSON(&okRes)
}

// 更新按钮
func patchButtonByButtonId(ctx *fiber.Ctx) error {
	buttonObj := new(model.Button)
	if util.IsInvalidJson(ctx, buttonObj) {
		return nil
	}
	if util.IsInvalidData(ctx, buttonObj) {
		return nil
	}

	err := service.PatchButtonService(buttonObj)
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success}
	return ctx.JSON(&okRes)
}
