package providers

import (
	"github.com/Mobilizes/materi-be-alpro/config"
	authController "github.com/Mobilizes/materi-be-alpro/modules/auth/controller"
	authRepo "github.com/Mobilizes/materi-be-alpro/modules/auth/repository"
	authService "github.com/Mobilizes/materi-be-alpro/modules/auth/service"
	userController "github.com/Mobilizes/materi-be-alpro/modules/user/controller"
	"github.com/Mobilizes/materi-be-alpro/modules/user/repository"
	userService "github.com/Mobilizes/materi-be-alpro/modules/user/service"
	"github.com/Mobilizes/materi-be-alpro/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func(i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (authService.JWTService, error) {
		return authService.NewJWTService(), nil
	})

	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[authService.JWTService](injector, constants.JWTService)

	userRepository := repository.NewUserRepository(db)
	refreshTokenRepository := authRepo.NewRefreshTokenRepository(db)

	userService := userService.NewUserService(userRepository, db)
	authService := authService.NewAuthService(userRepository, refreshTokenRepository, jwtService, db)

	do.Provide(
		injector, func(i *do.Injector) (userController.UserController, error) {
			return userController.NewUserController(i, userService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (authController.AuthController, error) {
			return authController.NewAuthController(i, authService), nil
		},
	)
}
