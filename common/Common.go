package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/global"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/object"
)

func PrintAndLog(msg string) {
	fmt.Println(msg)
	if global.SysConfig.Total.IsDebug {
		err := go_tool.Log(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func PrintOrLog(msg string) {
	if global.SysConfig.Total.IsDebug {
		err := go_tool.Log(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(msg)
	}
}

func GetSysConfig(fileName string) (*object.SysConfig, error) {
	path, err := go_tool.GetCurrPath()
	if err != nil {
		return nil, err
	}
	var config object.SysConfig
	_, err = toml.DecodeFile(path+"\\"+fileName, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func RefreshConfig(config *object.SysConfig) error {
	return nil
}

func rabbitMqInit() error {
	conn, err := global.RabbitMQ.GetConn()
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()

	err = global.RabbitMQ.QueueDeclareSimple(conn, "TktCreateYwdetail")
	if err != nil {
		return err
	}

	err = global.RabbitMQ.QueueBind(conn, "TktCreateYwdetail", "", "amq.fanout", true)
	if err != nil {
		return err
	}

	err = global.RabbitMQ.AddProducer("")
	if err != nil {
		return err
	}

	return nil
}