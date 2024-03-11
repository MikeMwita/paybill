package server

import (
	"github.com/MikeMwita/paybill/adapters"
)

const BaseUrl = "/api"

type Config struct {
	AuthUsecase adapters.AuthAdapter
}

//func NewServer(authUseCase adapters.AuthAdapter, cfg config.Config) *gin.Engine {
//	r := gin.Default()
//
//	// Initialize your handlers with necessary dependencies
//	myHandlers := handlers.NewMyHandlers(authUseCase)
//
//	options := handlers.GinServerOptions{
//		BaseURL: BaseUrl,
//	}
//
//	// Register handlers with options
//	handlers.RegisterHandlersWithOptions(r, myHandlers, options)
//
//	return r
//}
