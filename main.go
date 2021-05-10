package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "yinliuService/db"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("服务器启动失败")
	}
}
