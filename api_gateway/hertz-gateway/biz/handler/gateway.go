package handler

import (
	"bytes"
	"context"
	//"fmt"
	"net/http"

	"api_gateway/hertz-gateway/biz/errors"
	"api_gateway/hertz-gateway/biz/types"
	"api_gateway/kitex_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/kerrors"
	//"github.com/hertz-contrib/registry/nacos"
	//"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	//"log"
	"encoding/json"
)

type requiredParams struct {
	//Method     string `form:"method,required" json:"method"`
	//MerchantId string `form:"merchant_id,required" json:"merchant_id"`
	BizParams  string `form:"biz_params,required" json:"biz_params"`
}


var SvcMap = make(map[string]genericclient.Client)

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {
	//svcName := c.Param("action")
	svcName := c.Param("svc")
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, errors.New(common.Err_BadRequest))
		return
	}
	//var params requiredParams
	c.JSON(http.StatusOK, 666)
	//var params YourRequest
	//if err := c.BindAndValidate(&params); err != nil {
		//hlog.Error(err)
		//c.JSON(http.StatusOK, errors.New(common.Err_ServerMethodNotFound))
		//c.JSON(http.StatusOK, 1)
		//return
	//}
	
	//req, err := http.NewRequest(http.MethodPost, "", "")
//ceshi json
	body :=map[string]interface{}{
    		"message": "text",
	}
	//body :=map[string]interface{}{
    	//	"first": 1,
	//	"scond": 2,
	//}
	data, err := json.Marshal(body)
    	if err != nil {
        	panic(err)
    	}
	//url := "http://127.0.0.1:8888/DEFAULT_GROUP/hello"
	//url := "http://127.0.0.1:8888"
	req, err := http.NewRequest(http.MethodGet, "", bytes.NewBuffer(data))
	//req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(params.BizParams)))
	//if err != nil {
		//hlog.Warnf("new http request failed: %v", err)
		//c.JSON(http.StatusOK, errors.New(common.Err_RequestServerFail))
		//return
	//}
	//req.URL.Path = fmt.Sprintf("/%s/%s", svcName, params.Method)
	//req.URL.Path = fmt.Sprintf("/%s/%s", svcName, req.Method)

	//customReq, err := generic.FromHTTPRequest(c.Request)
	customReq, err := generic.FromHTTPRequest(req)
	//httpReq, err := http.NewRequest(c.Request.Method,"",c.Request.Body)
	//if err != nil {
    		// 处理错误
	//}
	//httpReq.Header = c.Request.Header
	//customReq, err := generic.FromHTTPRequest(httpReq)
	//customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		hlog.Errorf("convert request failed: %v", err)
		c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
		c.JSON(http.StatusOK, 2)
		return
	}

	resp, err := cli.GenericCall(ctx, "", customReq)
	respMap := make(map[string]interface{})
	if err != nil {
		hlog.Errorf("GenericCall err:%v", err)
		bizErr, ok := kerrors.FromBizStatusError(err)
		if !ok {
			c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
			c.JSON(http.StatusOK, 3)
			return
//
		}

		respMap[types.ResponseErrCode] = bizErr.BizStatusCode()
		respMap[types.ResponseErrMessage] = bizErr.BizMessage()
		c.JSON(http.StatusOK, respMap)
		return
	}
	realResp, ok := resp.(*generic.HTTPResponse)
	//realResp := "ok"
	if !ok {
		c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
		c.JSON(http.StatusOK, 4)
		return
	}
	c.JSON(http.StatusOK, 666)
	realResp.Body[types.ResponseErrCode] = 0
	realResp.Body[types.ResponseErrMessage] = "ok"
	c.JSON(http.StatusOK, realResp.Body)
	//realResp.Write(http.ResponseWriter w)
}






// Gateway handle the request with the query path of prefix `/gateway`.
func JGateway(ctx context.Context, c *app.RequestContext) {
	svcName := c.Param("svc")
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, errors.New(common.Err_BadRequest))
		return
	}
	//r, err := nacos.NewDefaultNacosResolver()
	//if err != nil {
		//log.Fatal(err)
		//return
	//}
	//cli.Use(sd.Discovery(r))


	//var params requiredParams
	c.JSON(http.StatusOK, 688)
	//req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(params.BizParams)))
	//if err != nil {
	//	hlog.Warnf("new http request failed: %v", err)
	//	c.JSON(http.StatusOK, errors.New(common.Err_RequestServerFail))
	//	return
	//}
	//req.URL.Path = fmt.Sprintf("/%s/%s", svcName, params.Method)
	//req.URL.Path = fmt.Sprintf("/%s/%s", svcName, req.Method)

	//customReq, err := generic.FromHTTPRequest(req)
	//if err != nil {
	//	hlog.Errorf("convert request failed: %v", err)
	//	c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
	//	c.JSON(http.StatusOK, 2)
	//	return
	//}

	resp, err := cli.GenericCall(ctx, "ExampleMethod", "{\"Msg\": \"hello\"}")
	//jsonResp, err := json.Marshal(resp)
	//if err != nil {
            //http.Error(w, "JSON encoding failed", http.StatusInternalServerError)
            //return
        //}

///
	//respMap := make(map[string]interface{})
	if err != nil {
		//hlog.Errorf("GenericCall err:%v", err)
		c.JSON(http.StatusOK, 3)
		//bizErr, ok := kerrors.FromBizStatusError(err)
		//if !ok {
			//c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
			//c.JSON(http.StatusOK, 3)
			//return
//
		//}

		//respMap[types.ResponseErrCode] = bizErr.BizStatusCode()
		//respMap[types.ResponseErrMessage] = bizErr.BizMessage()
		//c.JSON(http.StatusOK, respMap)
		//return
	}
	//realResp, ok := resp.(*generic.HTTPResponse)
	//realResp := "ok"
	//if !ok {
		//c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
		//c.JSON(http.StatusOK, 4)
		//return
	//}
	//c.JSON(http.StatusOK, 666)
	//realResp.Body[types.ResponseErrCode] = 0
	//realResp.Body[types.ResponseErrMessage] = "ok"
	c.JSON(http.StatusOK, resp)
	//realResp.Write(http.ResponseWriter w)
}
