package delivery

import (
	"web-app-crowdfounding/config"
	"web-app-crowdfounding/delivery/controller"
	"web-app-crowdfounding/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	engine         *gin.Engine
	useCaseManager manager.UseCaseManager
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)

	return &appServer{
		engine:         ginEngine,
		useCaseManager: usecase,
	}
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}

func (a *appServer) initHandlers() {
	userRoute := a.engine.Group("/api/v1")
	controller.NewUserController(userRoute, a.useCaseManager.UserUseCase())

}
