package route

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
	serviceBannerModule.Post("/new", middleware.TokenAuth, createNewBannerRoute)
	serviceBannerModule.Get("/list", middleware.TokenAuth, findBannersRoute)
	serviceBannerModule.Delete("/:id", middleware.TokenAuth, deleteBannerByBannerId)
	serviceBannerModule.Patch("/update", middleware.TokenAuth, patchBannerRoute)
	/****** ------ END -------------- **/

	/** ---------- app模块 ----------- **/
	//serviceAppModule := serviceModule.Group("/app")
	//serviceAppModule.Post("/new", middleware.TokenAuth, CreateNewAppRoute)
	/** ----------- END ------------- **/

	/** ----------- menu模块 --------- **/
	serviceMenuModule := serviceModule.Group("/menu")
	serviceMenuModule.Post("/new", middleware.TokenAuth, createNewMenuRoute)
	serviceMenuModule.Get("/list", middleware.TokenAuth, findAllMenusRoute)
	serviceMenuModule.Delete("/:id", middleware.TokenAuth, deleteMenuByMenuId)
	serviceMenuModule.Patch("/update", middleware.TokenAuth, patchMenuByMenuId)
	/** ----------- END ------------- **/

	/** ------------ button模块 -------- **/
	serviceButtonModule := serviceModule.Group("/button")
	serviceButtonModule.Post("/new", middleware.TokenAuth, createNewButtonRoute)
	serviceButtonModule.Get("/list", middleware.TokenAuth, findButtonsByAppId)
	serviceButtonModule.Delete("/:id", middleware.TokenAuth, deleteButtonByButtonId)
	serviceButtonModule.Patch("/:buttonId", middleware.TokenAuth, patchButtonByButtonId)
	/** ------------- END ------------- **/

	/** ------------- 前台的API数据 ----------- **/
	//apiModule := ylApiModule.Group("/front")
	//
	//apiModule.Get("/bannerList/:appId", FindBannersByAppId)
	//apiModule.Get("/MenuList/:appId", FindMenusByAppId)
	//apiModule.Get("/ButtonList/:appId", FindButtonsByAppId)

	/** ----------- 前台的API数据结尾 --------- **/
}
