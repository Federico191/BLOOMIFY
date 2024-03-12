package midtrans

//
//import (
//	"context"
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"github.com/midtrans/midtrans-go"
//	"github.com/midtrans/midtrans-go/coreapi"
//	"net/http"
//	"projectIntern/internal/entity"
//	"projectIntern/pkg/config"
//	"strconv"
//)
//
//type MidTransItf interface {
//	CreateTransaction(service entity.Service, booking entity.Booking) (*coreapi.ChargeResponse, error)
//	CallBackNotification(ctx context.Context)
//}
//
//type MidTrans struct {
//	client coreapi.Client
//	env    config.Env
//}
//
//func NewMidtrans(client coreapi.Client, env config.Env) MidTransItf {
//	return &MidTrans{client: client, env: env}
//}
//
//func (m MidTrans) CreateTransaction(service entity.Service, booking entity.Booking) (*coreapi.ChargeResponse, error) {
//	midtrans.ServerKey = m.env.ServerKey
//	midtrans.Environment = midtrans.Sandbox
//
//	chargeReq := &coreapi.ChargeReq{
//		PaymentType: coreapi.PaymentTypeBankTransfer,
//		TransactionDetails: midtrans.TransactionDetails{
//			OrderID:  booking.ID,
//			GrossAmt: int64(booking.GrossAmount),
//		},
//		Items: &[]midtrans.ItemDetails{
//			midtrans.ItemDetails{
//				ID:    strconv.Itoa(int(service.ID)),
//				Name:  service.Name,
//				Price: int64(service.Price),
//				Qty:   1,
//			},
//		},
//
//		BankTransfer: &coreapi.BankTransferDetails{
//			Bank: midtrans.BankBca,
//		},
//	}
//
//	resp, err := m.client.ChargeTransaction(chargeReq)
//	if err != nil {
//		return nil, err
//	}
//
//	return resp, nil
//}
//
//func (m MidTrans) CallBackNotification(ctx *gin.Context) {
//	// 1. Initialize empty map
//	var notificationPayload map[string]interface{}
//
//	// 2. Parse JSON request body and use it to set json to payload
//	err := json.NewDecoder(ctx.Request.Body).Decode(&notificationPayload)
//	if err != nil {
//		// do something on error when decode
//		return
//	}
//	// 3. Get order-id from payload
//	orderId, exists := notificationPayload["order_id"].(string)
//	if !exists {
//		// do something when key `order_id` not found
//		return
//	}
//
//	// 4. Check transaction to Midtrans with param orderId
//	transactionStatusResp, e := m.client.CheckTransaction(orderId)
//	if e != nil {
//		http.Error(ctx.Writer, e.GetMessage(), http.StatusInternalServerError)
//		return
//	} else {
//		if transactionStatusResp != nil {
//			// 5. Do set transaction status based on response from check transaction status
//			if transactionStatusResp.TransactionStatus == "capture" {
//				if transactionStatusResp.FraudStatus == "challenge" {
//					// TODO set transaction status on your database to 'challenge'
//					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
//				} else if transactionStatusResp.FraudStatus == "accept" {
//					// TODO set transaction status on your database to 'success'
//				}
//			} else if transactionStatusResp.TransactionStatus == "settlement" {
//				// TODO set transaction status on your databaase to 'success'
//			} else if transactionStatusResp.TransactionStatus == "deny" {
//				// TODO you can ignore 'deny', because most of the time it allows payment retries
//				// and later can become success
//			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
//				// TODO set transaction status on your databaase to 'failure'
//			} else if transactionStatusResp.TransactionStatus == "pending" {
//				// TODO set transaction status on your databaase to 'pending' / waiting payment
//			}
//		}
//	}
//	ctx.Writer.Header().Set("Content-Type", "application/json")
//	ctx.Writer.Write([]byte("ok"))
//}
//}
