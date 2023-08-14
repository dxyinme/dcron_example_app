package main

import (
	_ "app/docs"
	"app/routes"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			APP
//	@version		0.1
//	@description
//	@termsOfService	http://swagger.io/terms/
//
//	@BasePath /api/v1/user
//
//	@externalDocs.description		OpenAPI
//	@externalDocs.url						https://swagger.io/resources/open-api/
func main() {
	eng := routes.New()
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := eng.Run(); err != nil {
		log.Fatal(err)
	}
}
