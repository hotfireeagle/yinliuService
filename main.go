package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "yinliuService/db"
	"yinliuService/routes"
)

func main() {
	app := fiber.New()

	routes.InitRouter(app)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("服务器启动失败")
	}
}
