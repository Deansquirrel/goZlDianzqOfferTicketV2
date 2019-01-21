package object

import (
	"encoding/json"
	"fmt"
	"github.com/Deansquirrel/goZlDianzqOfferTicket/common"
	"github.com/Deansquirrel/goZlDianzqOfferTicket/global"
	"io/ioutil"
	"net/http"
)

type TicketNo struct {
}

func GetTktNoMulti(num int) ([]string, error) {
	rUrl := fmt.Sprintf("%s/Api/Number/GetTktNo_Multi?workerId=%d&nums=%d", global.SnoServer, global.SnoWorkerId, num)
	resp, err := http.Get(rUrl)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			common.MyLog(err.Error())
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	var r []string
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetSno(prefix string) (string, error) {
	rUrl := fmt.Sprintf("%s/Api/Number/GetSno?workerId=%d&&prefix=%s", global.SnoServer, global.SnoWorkerId, prefix)
	resp, err := http.Get(rUrl)
	if err != nil {
		return "", err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			common.MyLog(err.Error())
		}
	}()
	sno, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(sno), nil
}
