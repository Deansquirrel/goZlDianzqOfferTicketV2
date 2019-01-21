package yw

import (
	"github.com/kataras/iris"
)

func handler(ctx iris.Context) {
	//response := getResponse(ctx)
	//_, err := ctx.Write(getResponseData(response))
	//if err != nil {
	//	common.PrintAndLog(err.Error())
	//}
	return
}

//
//func getResponse(ctx iris.Context) (response yw.ResponseCreateLittleTkt) {
//	request, err := yw.GetRequestCreateLittleTktByContext(ctx)
//	if err != nil {
//		return getErrorResponse(request, ctx, err)
//	}
//	err = request.CheckRequest()
//	if err != nil {
//		return getErrorResponse(request, ctx, err)
//	}
//	return yw.GetResponseCreateLittleTkt(ctx, &request)
//}
//
//func getResponseData(response yw.ResponseCreateLittleTkt) []byte {
//	data, err := json.Marshal(response)
//	if err != nil {
//		common.PrintAndLog(err.Error())
//		return []byte(err.Error())
//	} else {
//		return data
//	}
//}
//
//func getErrorResponse(request yw.RequestCreateLittleTkt, ctx iris.Context, err error) (response yw.ResponseCreateLittleTkt) {
//	response = yw.GetResponseCreateLittleTktError(&request, err, ctx.GetStatusCode())
//	return response
//}
