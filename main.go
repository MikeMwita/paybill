package main

import (
	"fmt"
	"net/http"
)

// @title Paybill API
// @version 1.0
// @description This is a simple Paybill API
// @contact.name API Support
// @termsOfService demo.com
// @contact.url http://demo.com/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api
// @Schemes http https
// @query.collection.format multi
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println(fmt.Sprintf("Number of ytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)

	//userService := service.NewUserService()
	//myHandlers := handlers.NewMyHandlers(userService)
	//
	//// Initialize a Gin router
	//router := gin.Default()
	//
	//// Call the RegisterHandlers function from your generated code
	//handlers.RegisterHandlers(router, myHandlers)
	//
	//// Start your server
	//router.Run(":8080")
}
