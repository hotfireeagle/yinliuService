package routes

import (
	"github.com/gofiber/fiber/v2"
	"yinliuService/middleware"
)

func InitRouter(app *fiber.App) {
	ylApiModule := app.Group("/ylApi")

	serviceModule := ylApiModule.Group("/ope")

	serviceUserModule := serviceModule.Group("/user")
	serviceUserModule.Post("/login", OpsUserDoLogin)
	serviceUserModule.Post("/create_shadow_not_public12131213", OpsUserDoCreate)

	/** -------- banner模块 ---------- **/
	serviceBannerModule := serviceModule.Group("/banner")
	serviceBannerModule.Post("/new", middleware.TokenAuth, CreateNewBannerRoute)
	serviceBannerModule.Get("/list/:appId", middleware.TokenAuth, FindBannersByAppId)
	serviceBannerModule.Delete("/:bannerId", middleware.TokenAuth, DeleteBannerByBannerId)
	serviceBannerModule.Patch("/:bannerId", middleware.TokenAuth, PatchBannerByBannerId)
	/****** ------ END -------------- **/

	/** ---------- app模块 ----------- **/
	serviceAppModule := serviceModule.Group("/app")
	serviceAppModule.Post("/new", middleware.TokenAuth, CreateNewAppRoute)
	/** ----------- END ------------- **/

	/** ----------- menu模块 --------- **/
	serviceMenuModule := serviceModule.Group("/menu")
	serviceMenuModule.Post("/new", middleware.TokenAuth, CreateNewMenuRoute)
	serviceMenuModule.Get("/list/:appId", middleware.TokenAuth, FindMenusByAppId)
	serviceMenuModule.Delete("/:menuId", middleware.TokenAuth, DeleteMenuByMenuId)
	serviceMenuModule.Patch("/:menuId", middleware.TokenAuth, PatchMenuByMenuId)
	/** ----------- END ------------- **/
}
