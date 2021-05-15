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

	/** ------------ button模块 -------- **/
	serviceButtonModule := serviceModule.Group("/button")
	serviceButtonModule.Post("/new", middleware.TokenAuth, CreateNewButtonRoute)
	serviceButtonModule.Get("/list/:appId", middleware.TokenAuth, FindButtonsByAppId)
	serviceButtonModule.Delete("/:buttonId", middleware.TokenAuth, DeleteButtonByButtonId)
	serviceButtonModule.Patch("/:buttonId", middleware.TokenAuth, PatchButtonByButtonId)
	/** ------------- END ------------- **/

	/** ------------- 前台的API数据 ----------- **/
	apiModule := ylApiModule.Group("/front")

	apiModule.Get("/bannerList/:appId", FindBannersByAppId)
	apiModule.Get("/MenuList/:appId", FindMenusByAppId)
	apiModule.Get("/ButtonList/:appId", FindButtonsByAppId)

	/** ----------- 前台的API数据结尾 --------- **/
}
