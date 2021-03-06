package repository

import (
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/common"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/global"
	"github.com/kataras/iris/core/errors"
)

type PeiZhRepository struct {
}

//获取配置库连接对象
func getPeiZhDbConn() error {
	if CheckV(global.PeiZhDbConn) {
		return nil
	}
	conn, err := GetDbConn(global.SysConfig.PeiZhDb.Server,
		global.SysConfig.PeiZhDb.Port,
		global.SysConfig.PeiZhDb.DbName,
		global.SysConfig.PeiZhDb.User,
		global.SysConfig.PeiZhDb.Password)
	if err != nil {
		global.PeiZhDbConn = nil
		return err
	}

	err = conn.Ping()
	if err != nil {
		global.PeiZhDbConn = nil
		return err
	}

	//conn.SetMaxIdleConns(30)
	//conn.SetMaxOpenConns(30)
	//conn.SetConnMaxLifetime(time.Second * 60 * 10)
	global.PeiZhDbConn = conn

	return nil
}

//从xtwxappidjoininfo获取配置
func (pzR *PeiZhRepository) GetXtWxAppIdJoinInfo(jPeiZh string, jKey string, jIsForbid int) (string, error) {
	if !CheckV(global.PeiZhDbConn) {
		err := getPeiZhDbConn()
		if err != nil {
			return "", err
		}
	}

	conn := global.PeiZhDbConn
	//defer func() {
	//	errLs := conn.Close()
	//	if errLs != nil {
	//		common.PrintOrLog(errLs.Error())
	//	}
	//}()

	stmt, err := conn.Prepare("" +
		"SELECT jvalue FROM xtwxappidjoininfo " +
		"WHERE jpeizh = ? and jkey = ? AND JISFORBID = ?")
	if err != nil {
		return "", err
	}
	defer func() {
		errLs := stmt.Close()
		if errLs != nil {
			common.PrintOrLog(errLs.Error())
		}
	}()

	rows, err := stmt.Query(jPeiZh, jKey, jIsForbid)
	if err != nil {
		return "", err
	}
	defer func() {
		errLs := rows.Close()
		if errLs != nil {
			common.PrintOrLog(errLs.Error())
		}
	}()

	var valList []string
	for rows.Next() {
		var val string
		err := rows.Scan(&val)
		if err != nil {
			return "", err
		}
		valList = append(valList, val)
	}

	if len(valList) > 0 {
		return valList[0], nil
	} else {
		return "", errors.New("未获取到配置值")
	}
}

//从xtmappingdbconn获取连接信息
func (pzR *PeiZhRepository) GetXtMappingDbConnInfo(appId string, miKvName string, miIdType string) ([]dbConnInfo, error) {
	if !CheckV(global.PeiZhDbConn) {
		err := getPeiZhDbConn()
		if err != nil {
			return nil, err
		}
	}

	conn := global.PeiZhDbConn
	//defer func() {
	//	errLs := conn.Close()
	//	if errLs != nil {
	//		common.PrintOrLog(errLs.Error())
	//	}
	//}()

	stmt, err := conn.Prepare("" +
		"select miid,mconnstr " +
		"from xtmappingdbconn where appid = ? and miidtype = ? and mikvname = ?")
	if err != nil {
		return nil, err
	}
	defer func() {
		errLs := stmt.Close()
		if errLs != nil {
			common.PrintOrLog(errLs.Error())
		}
	}()

	rows, err := stmt.Query(appId, miIdType, miKvName)
	if err != nil {
		return nil, err
	}
	defer func() {
		errLs := rows.Close()
		if errLs != nil {
			common.PrintOrLog(errLs.Error())
		}
	}()

	dbConnInfoList := make([]dbConnInfo, 0)

	for rows.Next() {
		var val dbConnInfo
		err := rows.Scan(&val.MiId, &val.MConnStr)
		if err != nil {
			return nil, err
		}
		dbConnInfoList = append(dbConnInfoList, val)
	}

	return dbConnInfoList, nil
}

type dbConnInfo struct {
	MiId     int
	MConnStr string
}
