package handlers

import (
	"github.com/MikeMwita/paybill/internal/routes/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handlers *handlers.MyHandlers) {

	authApi := router.Group("/api/auth")
	{
		authApi.POST("/register", handlers.PostApiAuthRegister)
		authApi.POST("/login", handlers.PostApiAuthLogin)
		authApi.POST("/logout", handlers.PostApiAuthLogout)
		authApi.POST("/refresh", handlers.PostApiAuthRefresh)
		authApi.POST("/reset", handlers.PostApiAuthReset)
		authApi.PUT("/reset/:token", handlers.PutApiAuthResetToken)
		authApi.POST("/verify", handlers.PostApiAuthVerify)
		authApi.GET("/verify/:token", handlers.GetApiAuthVerifyToken)
		authApi.GET("/profile", handlers.GetApiAuthProfile)
		authApi.PUT("/profile", handlers.PutApiAuthProfile)
	}

}
