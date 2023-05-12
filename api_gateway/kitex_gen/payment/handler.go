package main

import (
	"context"
	payment "api_gateway/kitex_gen/payment/kitex_gen/payment"
)

// PaymentSvcImpl implements the last service interface defined in the IDL.
type PaymentSvcImpl struct{}

// UnifyPay implements the PaymentSvcImpl interface.
func (s *PaymentSvcImpl) UnifyPay(ctx context.Context, req *payment.UnifyPayReq) (resp *payment.UnifyPayResp, err error) {
	// TODO: Your code here...
	return
}

// QRPay implements the PaymentSvcImpl interface.
func (s *PaymentSvcImpl) QRPay(ctx context.Context, req *payment.QRPayReq) (resp *payment.QRPayResp, err error) {
	// TODO: Your code here...
	return
}

// QueryOrder implements the PaymentSvcImpl interface.
func (s *PaymentSvcImpl) QueryOrder(ctx context.Context, req *payment.QueryOrderReq) (resp *payment.QueryOrderResp, err error) {
	// TODO: Your code here...
	return
}

// CloseOrder implements the PaymentSvcImpl interface.
func (s *PaymentSvcImpl) CloseOrder(ctx context.Context, req *payment.CloseOrderReq) (resp *payment.CloseOrderResp, err error) {
	// TODO: Your code here...
	return
}
