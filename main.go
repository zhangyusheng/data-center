package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zhangyusheng/data-center/models"
	"github.com/zhangyusheng/data-center/pkg/gredis"
	"github.com/zhangyusheng/data-center/pkg/logging"
	"github.com/zhangyusheng/data-center/pkg/setting"
	"github.com/zhangyusheng/data-center/pkg/util"
	"github.com/zhangyusheng/data-center/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/zhangyusheng/data-center
// @license.name MIT
// @license.url https://github.com/zhangyusheng/data-center/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
