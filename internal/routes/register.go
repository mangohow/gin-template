package routes

import (
	"github.com/mangohow/easygin"
	"github.com/mangohow/gin-template/internal/controller"
)

func Register(r *easygin.EasyGin) {
	helloController := controller.NewHelloController()
	{
		group := r.Group("/api")
		group.GET("/hello", helloController.Hello)
	}
}
