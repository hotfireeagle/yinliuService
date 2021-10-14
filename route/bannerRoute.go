package route

import (
	"github.com/gofiber/fiber/v2"
	"yinliuService/model"
	"yinliuService/service"
	"yinliuService/util"
)

// 新建banner
func createNewBannerRoute(ctx *fiber.Ctx) error {
	createBannerVO := new(model.Banner)
	if util.IsInvalidJson(ctx, createBannerVO) {
		return nil
	}
	if util.IsInvalidData(ctx, createBannerVO) {
		return nil
	}

	createBannerVO.GenerateUUID()
	if err := service.CreateBannerService(createBannerVO); err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	} else {
		okRes := model.IResponse{Status: model.Success}
		return ctx.JSON(&okRes)
	}
}

// 查找出所有的banner数据
func findBannersRoute(ctx *fiber.Ctx) error {
	res, err := service.SearchAllBannerService()
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success, Data: res}
	return ctx.JSON(&okRes)
}

/**
** 删除banner
 */
func deleteBannerByBannerId(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := service.DelBannerService(id)
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success}
	return ctx.JSON(&okRes)
}

// 更新banner
func patchBannerRoute(ctx *fiber.Ctx) error {
	banner := new(model.Banner)
	if util.IsInvalidJson(ctx, banner) {
		return nil
	}
	if util.IsInvalidData(ctx, banner) {
		return nil
	}

	err := service.UpdateBannerService(banner)
	if err != nil {
		errRes := model.IResponse{Status: model.Err, Msg: "系统异常", ErrLog: err.Error()}
		return ctx.JSON(&errRes)
	}
	okRes := model.IResponse{Status: model.Success}
	return ctx.JSON(&okRes)
}
