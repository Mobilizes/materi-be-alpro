package user

import (
	"github.com/Mobilizes/materi-be-alpro/middlewares"
	"github.com/Mobilizes/materi-be-alpro/modules/auth/service"
	"github.com/Mobilizes/materi-be-alpro/modules/user/controller"
	"github.com/Mobilizes/materi-be-alpro/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	userController := do.MustInvoke[controller.UserController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	userRoutes := server.Group("/api/user")
	{
		userRoutes.GET("", userController.GetAllUser)
		userRoutes.GET("/me", middlewares.Authenticate(jwtService), userController.Me)
		userRoutes.PUT("/:id", middlewares.Authenticate(jwtService), userController.Update)
		userRoutes.DELETE("/:id", middlewares.Authenticate(jwtService), userController.Delete)
	}
}
