// Code generated by Kitex v0.4.4. DO NOT EDIT.

package paymentsvc

import (
	"context"
	payment "api_gateway/kitex_gen/payment/kitex_gen/payment"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return paymentSvcServiceInfo
}

var paymentSvcServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PaymentSvc"
	handlerType := (*payment.PaymentSvc)(nil)
	methods := map[string]kitex.MethodInfo{
		"UnifyPay":   kitex.NewMethodInfo(unifyPayHandler, newPaymentSvcUnifyPayArgs, newPaymentSvcUnifyPayResult, false),
		"QRPay":      kitex.NewMethodInfo(qRPayHandler, newPaymentSvcQRPayArgs, newPaymentSvcQRPayResult, false),
		"QueryOrder": kitex.NewMethodInfo(queryOrderHandler, newPaymentSvcQueryOrderArgs, newPaymentSvcQueryOrderResult, false),
		"CloseOrder": kitex.NewMethodInfo(closeOrderHandler, newPaymentSvcCloseOrderArgs, newPaymentSvcCloseOrderResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "payment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func unifyPayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment.PaymentSvcUnifyPayArgs)
	realResult := result.(*payment.PaymentSvcUnifyPayResult)
	success, err := handler.(payment.PaymentSvc).UnifyPay(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentSvcUnifyPayArgs() interface{} {
	return payment.NewPaymentSvcUnifyPayArgs()
}

func newPaymentSvcUnifyPayResult() interface{} {
	return payment.NewPaymentSvcUnifyPayResult()
}

func qRPayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment.PaymentSvcQRPayArgs)
	realResult := result.(*payment.PaymentSvcQRPayResult)
	success, err := handler.(payment.PaymentSvc).QRPay(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentSvcQRPayArgs() interface{} {
	return payment.NewPaymentSvcQRPayArgs()
}

func newPaymentSvcQRPayResult() interface{} {
	return payment.NewPaymentSvcQRPayResult()
}

func queryOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment.PaymentSvcQueryOrderArgs)
	realResult := result.(*payment.PaymentSvcQueryOrderResult)
	success, err := handler.(payment.PaymentSvc).QueryOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentSvcQueryOrderArgs() interface{} {
	return payment.NewPaymentSvcQueryOrderArgs()
}

func newPaymentSvcQueryOrderResult() interface{} {
	return payment.NewPaymentSvcQueryOrderResult()
}

func closeOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment.PaymentSvcCloseOrderArgs)
	realResult := result.(*payment.PaymentSvcCloseOrderResult)
	success, err := handler.(payment.PaymentSvc).CloseOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentSvcCloseOrderArgs() interface{} {
	return payment.NewPaymentSvcCloseOrderArgs()
}

func newPaymentSvcCloseOrderResult() interface{} {
	return payment.NewPaymentSvcCloseOrderResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UnifyPay(ctx context.Context, req *payment.UnifyPayReq) (r *payment.UnifyPayResp, err error) {
	var _args payment.PaymentSvcUnifyPayArgs
	_args.Req = req
	var _result payment.PaymentSvcUnifyPayResult
	if err = p.c.Call(ctx, "UnifyPay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QRPay(ctx context.Context, req *payment.QRPayReq) (r *payment.QRPayResp, err error) {
	var _args payment.PaymentSvcQRPayArgs
	_args.Req = req
	var _result payment.PaymentSvcQRPayResult
	if err = p.c.Call(ctx, "QRPay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryOrder(ctx context.Context, req *payment.QueryOrderReq) (r *payment.QueryOrderResp, err error) {
	var _args payment.PaymentSvcQueryOrderArgs
	_args.Req = req
	var _result payment.PaymentSvcQueryOrderResult
	if err = p.c.Call(ctx, "QueryOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CloseOrder(ctx context.Context, req *payment.CloseOrderReq) (r *payment.CloseOrderResp, err error) {
	var _args payment.PaymentSvcCloseOrderArgs
	_args.Req = req
	var _result payment.PaymentSvcCloseOrderResult
	if err = p.c.Call(ctx, "CloseOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
