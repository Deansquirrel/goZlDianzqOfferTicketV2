package yw

import (
	"encoding/json"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV2/common"
	"github.com/kataras/iris"
)

func handler(ctx iris.Context) {
	response := getResponse(ctx)
	_, err := ctx.Write(getResponseData(response))
	if err != nil {
		common.PrintAndLog(err.Error())
	}
	return
}

func getResponse(ctx iris.Context) (response ResponseCreateLittleTkt) {
	request, err := GetRequestCreateLittleTktByContext(ctx)
	if err != nil {
		return getErrorResponse(request, ctx, err)
	}
	err = request.CheckRequest()
	if err != nil {
		return getErrorResponse(request, ctx, err)
	}
	return GetResponseCreateLittleTkt(ctx, &request)
}

func getResponseData(response ResponseCreateLittleTkt) []byte {
	data, err := json.Marshal(response)
	if err != nil {
		common.PrintAndLog(err.Error())
		return []byte(err.Error())
	} else {
		return data
	}
}

func getErrorResponse(request RequestCreateLittleTkt, ctx iris.Context, err error) (response ResponseCreateLittleTkt) {
	response = GetResponseCreateLittleTktError(&request, err, ctx.GetStatusCode())
	return response
}
