package handler

import (
	"context"
	"net/http"
	"github.com/zz1219/api_gateway/hertz-gateway/biz/errors"
	"github.com/zz1219/api_gateway/kitex_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client/genericclient"
	"strings"
)
var SvcMap = make(map[string]genericclient.Client)

// Gateway handle the request with the query path of prefix `/gateway`.
func JGateway(ctx context.Context, c *app.RequestContext) {
	sv := string(c.Request.Path())
	parts := strings.Split(sv, "/")
	svcName := parts[1]
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, errors.New(common.Err_BadRequest))
		return
	}
	action := c.Param("action")
	if len(action) == 0{
		c.JSON(http.StatusOK, errors.New(common.Err_ServerMethodNotFound))
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
