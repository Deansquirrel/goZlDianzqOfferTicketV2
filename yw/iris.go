package yw

import (
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/common"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/global"
	"github.com/kataras/iris"
	"strconv"
)

func StartWebServer() {
	app := iris.New()
	app.Post("/", handler)
	addr := ":" + strconv.Itoa(global.SysConfig.Total.Port)
	err := app.Run(iris.Addr(addr), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
	if err != nil {
		common.PrintAndLog(err.Error())
	}
}
