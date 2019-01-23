package global

import (
	"database/sql"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/object"
)

var SysConfig *object.SysConfig
var RabbitMQ *go_tool.MyRabbitMQ
var Redis *go_tool.MyRedis

var PeiZhDbConn *sql.DB
var HxDbConnMap map[string]*sql.DB

var RedisDbId1 int
var RedisDbId2 int

var SnoServer string
var SnoWorkerId int
