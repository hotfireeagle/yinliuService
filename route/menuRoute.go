package route

import (
	"github.com/gofiber/fiber/v2"
	"yinliuService/model"
	"yinliuService/service"
	"yinliuService/tool"
	"yinliuService/util"
)

// 新增菜单栏
func createNewMenuRoute(ctx *fiber.Ctx) error {
	menu := new(model.Menu)
	if util.IsInvalidJson(ctx, menu) {
		return nil
	}
	if util.IsInvalidData(ctx, menu) {
		return nil
	}

	menu.GenerateUUID()
	err := service.CreateNewMenuService(menu)
	if err != nil {
		tool.ErrLog(err)
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success}
	return ctx.JSON(&okRes)
}

// 查找出所有的菜单栏
func findAllMenusRoute(ctx *fiber.Ctx) error {
	result, err := service.FindMenusService()
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success, Data: *result}
	return ctx.JSON(&okRes)
}

// 删除菜单
func deleteMenuByMenuId(ctx *fiber.Ctx) error {
	menuId := ctx.Params("id")
	err := service.DeleteMenuService(menuId)
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success}
	return ctx.JSON(&okRes)
}

// 更新menu
func patchMenuByMenuId(ctx *fiber.Ctx) error {
	menu := new(model.Menu)
	if util.IsInvalidJson(ctx, menu) {
		return nil
	}
	if util.IsInvalidData(ctx, menu) {
		return nil
	}

	err := service.PatchMenuService(menu)
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success}
	return ctx.JSON(&okRes)
}
