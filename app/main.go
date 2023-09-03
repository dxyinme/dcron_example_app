package main

import (
	"app/config"
	_ "app/docs"
	"app/internal/crontasks"
	"app/internal/customerdb"
	"app/internal/db"
	"app/routes"
	"flag"
	"log"

	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	configFile = flag.String("f", "etc/app.yaml", "config file")
)

//	@title			APP
//	@version		0.1
//	@description
//	@termsOfService	http://swagger.io/terms/
//
//	@BasePath /api/v1
//
//	@externalDocs.description		OpenAPI
//	@externalDocs.url						https://swagger.io/resources/open-api/
func main() {
	flag.Parse()
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	err := config.LoadConfig(*configFile)
	if err != nil {
		panic(err)
	}

	db.SelfStoreUtil{}.Initial()
	customerdb.DBStoresUtil{}.Initial()
	crontasks.CronTasksContainerUtil{}.Initial()

	eng := routes.New()
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := eng.Run(); err != nil {
		log.Fatal(err)
	}
}
