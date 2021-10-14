package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "yinliuService/environmentVariable"
	"yinliuService/route"
)

func main() {
	app := fiber.New()
	// 开启cors配置
	app.Use(cors.New())

	route.InitRouter(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
