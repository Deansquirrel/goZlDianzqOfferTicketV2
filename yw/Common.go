package yw

import (
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/common"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/global"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/object"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/repository"
	"github.com/kataras/iris/core/errors"
	"strconv"
	"strings"
	"time"
)

func RefreshConfig(config *object.SysConfig) error {
	if config == nil {
		return errors.New("config对象不能为空")
	}
	//==================================================================================================================
	sConfig, err := global.SysConfig.GetConfigStr()
	if err != nil {
		common.PrintAndLog(err.Error())
	} else {
		common.PrintAndLog(sConfig)
	}
	//==================================================================================================================
	pZhR := repository.PeiZhRepository{}

	//获取Hx库连接信息

	//获取Redis连接信息
	redisConfigStr, err := pZhR.GetXtWxAppIdJoinInfo(global.SysConfig.Total.JPeiZh, "SERedis", 0)
	if err != nil {
		return err
	}
	redisConfig := strings.Split(redisConfigStr, "|")
	if len(redisConfig) != 2 {
		return errors.New("redis配置参数异常.expected 2 , got " + strconv.Itoa(len(redisConfig)))
	}

	global.Redis = go_tool.NewRedis(redisConfig[0], redisConfig[1], 5000, 5000, 5)
	if err != nil {
		return err
	}

	redisDbId1Str, err := pZhR.GetXtWxAppIdJoinInfo(global.SysConfig.Total.JPeiZh, "RedisDbId1", 0)
	if err != nil {
		return err
	}
	global.RedisDbId1, err = strconv.Atoi(redisDbId1Str)
	if err != nil {
		return err
	}

	redisDbId2Str, err := pZhR.GetXtWxAppIdJoinInfo(global.SysConfig.Total.JPeiZh, "RedisDbId2", 0)
	if err != nil {
		return err
	}
	global.RedisDbId2, err = strconv.Atoi(redisDbId2Str)
	if err != nil {
		return err
	}

	//获取RabbitMQ连接信息
	rabbitMQConfigStr, err := pZhR.GetXtWxAppIdJoinInfo(global.SysConfig.Total.JPeiZh, "RabbitConnection", 0)
	if err != nil {
		return err
	}
	rabbitMQConfig := strings.Split(rabbitMQConfigStr, "|")
	if len(rabbitMQConfig) != 5 {
		return errors.New("rabbitMQ配置参数异常.expected 5 , got " + strconv.Itoa(len(rabbitMQConfig)))
	}
	rabbitMQPort, err := strconv.Atoi(rabbitMQConfig[1])
	if err != nil {
		return err
	}
	global.RabbitMQ = go_tool.NewRabbitMQ(rabbitMQConfig[3], rabbitMQConfig[4], rabbitMQConfig[0], rabbitMQPort, rabbitMQConfig[2], time.Second*60, time.Millisecond*500, 3, time.Second*5)

	//获取SnoServer信息
	global.SnoServer, err = pZhR.GetXtWxAppIdJoinInfo(global.SysConfig.Total.JPeiZh, "SnoServer", 0)
	if err != nil {
		return err
	}
	snoWorkIdStr, err := pZhR.GetXtWxAppIdJoinInfo(global.SysConfig.Total.JPeiZh, "WorkerId", 0)
	if err != nil {
		return err
	}
	global.SnoWorkerId, err = strconv.Atoi(snoWorkIdStr)
	if err != nil {
		return err
	}
	global.SnoWorkerId, err = strconv.Atoi(snoWorkIdStr)
	if err != nil {
		return err
	}

	err = rabbitMqInit()
	if err != nil {
		return err
	}

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

	//err = global.RabbitMQ.AddConsumer("","TktCreateYwdetail",lsHandler)

	return nil
}

//func lsHandler(msg string){
//	common.PrintOrLog(msg)
//}
//
