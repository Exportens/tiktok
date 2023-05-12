
namespace go payment

struct UnifyPayReq {
    1: string out_order_no,
    2: i64 total_amount,
    3: string subject,
    4: string merchant_id,
    5: string pay_way,
    6: string app_id,
    7: string sub_open_id,
    8: string notify_url,
    9: string client_ip,
    10: i32 order_expiration
}

struct UnifyPayResp {
    1: string merchant_id,
    2: string sub_merchant_id,
    3: string out_order_no,
    4: string jspay_info,
    5: string pay_way,
}

struct QRPayReq {
    1: string out_order_no,
    2: i64 total_amount,
    3: string subject,
    4: string merchant_id,
    5: string auth_code,
    6: string notify_url,
    7: string client_ip,
}

struct QRPayResp {
    1: string merchant_id,
    2: string sub_merchant_id,
    3: string out_order_no,
    4: i8 order_status,
    5: string pay_way,
    6: string open_id,
    7: string out_transaction_id,
    8: string sub_openid,
}

struct QueryOrderReq {
    1: string out_order_no,
}

struct QueryOrderResp {
    1: i8 order_status,
}

struct CloseOrderReq {
     1: string out_order_no,
}

struct CloseOrderResp {
}

service PaymentSvc {
    UnifyPayResp UnifyPay(1: UnifyPayReq req)( api.post = '/payment/unifypay', api.param = 'true')

    QRPayResp QRPay(1: QRPayReq req)( api.post = '/payment/qrpay', api.param = 'true')

    QueryOrderResp QueryOrder(1: QueryOrderReq req)( api.post = '/payment/queryorder', api.param = 'true')

    CloseOrderResp CloseOrder(1: CloseOrderReq req)( api.post = '/payment/closeorder', api.param = 'true')
}
