package routes

import (
	"golang_project/controllers"
	"golang_project/docs"
	"golang_project/repositories"
	"golang_project/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	friendConnectionRepository repositories.FriendConnectionRepository = repositories.New()
	friendConnectionService    services.FriendConnectionService        = services.New(friendConnectionRepository)
	friendConnectionController controllers.FriendConnectionController  = controllers.New(friendConnectionService)
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//1. Done
			v1.POST("/friends/createConnection", friendConnectionController.CreateFriendConnection)

			//2. Done
			v1.POST("/friends/showFriendsByEmail", friendConnectionController.GetFriendListByEmail)

			//3. Done
			v1.POST("/friends/showCommonFriendList", friendConnectionController.ShowCommonFriendList)

			//4. Done
			v1.POST("/friends/subscribeFromEmail", friendConnectionController.SubscribeFromEmail)

			//5. Done
			v1.POST("/friends/blockSubscribeByEmail", friendConnectionController.BlockSuscribeByEmail)

			//6. Done
			v1.POST("/friends/showSubscribingEmailListByEmail", friendConnectionController.GetSubscribingEmailListByEmail)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
