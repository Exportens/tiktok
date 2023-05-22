package handler

import (
	"context"
	"net/http"
	"api_gateway/hertz-gateway/biz/errors"
	"api_gateway/kitex_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client/genericclient"
)
var SvcMap = make(map[string]genericclient.Client)

// Gateway handle the request with the query path of prefix `/gateway`.
func JGateway(ctx context.Context, c *app.RequestContext) {
	svcName := c.Param("svc")
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, errors.New(common.Err_BadRequest))
		return
	}
	action := c.Param("action")
	if len(action) == 0{
		c.JSON(http.StatusOK, errors.New(common.Err_RequestServerFail))
		return
	}
	data := c.Request.Body()
	str := string(data)
	resp, err := cli.GenericCall(ctx, action,str )
	if err != nil {
		c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
	}
	c.JSON(http.StatusOK, "996ers")
	c.JSON(http.StatusOK, resp)
}
