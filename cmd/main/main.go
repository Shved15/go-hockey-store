package main

import (
	"github.com/gin-gonic/gin"
	"go-code/go-hockey-shop/handlers"
)

func main() {
	// создание экземпляра Gin
	r := gin.Default()

	// обработчик для статических файлов
	r.Static("/static", "./static")

	// регистрация маршрутов
	r.GET("/", handlers.IndexView)

	// запуск сервера
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
