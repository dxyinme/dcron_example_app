package main

import (
	"app/config"
	"app/controllers/helper"
	_ "app/docs"
	"app/internal/common/innercall"
	"app/internal/crontasks"
	"app/internal/customerdb"
	"app/internal/db"
	"app/routes"
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	configFile = flag.String("f", "etc/app.yaml", "config file")
	configMode = flag.String("mode", "FromFile", "config mode [ FromFile | FromEnv ]")
)

// @title												APP
// @version											0.1
// @description
// @termsOfService							http://swagger.io/terms/
//
// @BasePath										/api/v1
//
// @externalDocs.description		OpenAPI
// @externalDocs.url						https://swagger.io/resources/open-api/
func main() {
	flag.Parse()
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	if err := config.LoadConfig(*configMode, *configFile); err != nil {
		panic(err)
	}
	logrus.Debugf("config:%v", config.I())

	db.SelfStoreUtil{}.Initial()
	customerdb.DBStoresUtil{}.Initial()
	crontasks.CronTasksContainerUtil{}.Initial()
	innercall.InnerCallUtil{}.Initial()
	go helper.InnerCallLoop()

	eng := routes.New()
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := eng.Run(
		fmt.Sprintf(":%d", config.I().Port),
	); err != nil {
		logrus.Fatal(err)
	}
}
