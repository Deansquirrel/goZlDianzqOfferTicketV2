package main

import (
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/common"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/global"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/yw"
)

func main() {
	//==================================================================================================================
	config, err := common.GetSysConfig("config.toml")
	if err != nil {
		common.PrintAndLog("加载配置文件时遇到错误：" + err.Error())
		return
	}
	global.SysConfig = config
	err = yw.RefreshConfig(global.SysConfig)
	if err != nil {
		common.PrintAndLog("刷新配置时遇到错误：" + err.Error())
		return
	}
	//==================================================================================================================

	//==================================================================================================================
	common.PrintOrLog("程序启动")
	defer common.PrintOrLog("程序退出")
	//==================================================================================================================
	yw.StartWebServer()
}
