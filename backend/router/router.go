package router

import (
	"blog/api/v1"
	"blog/config"
	"blog/middleware"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

func InitRouter() {
	gin.SetMode(config.ServerMode)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("UsernameValidator", utils.UsernameValidator)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("TelValidator", utils.TelValidator)
	}
	router := gin.Default()
	router.Use(middleware.LogMiddleWare())
	router.POST("/api/v1/login", v1.Login)
	router.POST("/api/v1/register", v1.Register)
	routerGroup := router.Group("/api/v1")
	routerGroup.Use(middleware.JwtMiddleWare())
	{
		routerGroup.PUT("/user/:id", v1.UpdateUser)
		routerGroup.DELETE("/user/:id", v1.DeleteUser)
	}
	log.Fatal(router.Run(config.ServerPort))
}
