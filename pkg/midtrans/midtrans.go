package midtrans

import (
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"net/http"
	"projectIntern/internal/entity"
	"projectIntern/pkg/config"
	"projectIntern/pkg/response"
	"strconv"
)

type MidTransItf interface {
	CreateTransaction(service entity.Service, booking entity.Booking, user entity.User) (*coreapi.ChargeResponse, error)
	CheckTransactionStatus(ctx *gin.Context)
}

type MidTrans struct {
	client coreapi.Client
	env    config.Env
}

func NewMidtrans(client coreapi.Client, env config.Env) MidTransItf {
	return &MidTrans{client: client, env: env}
}

func (m MidTrans) CreateTransaction(service entity.Service, booking entity.Booking, user entity.User) (*coreapi.ChargeResponse, error) {
	m.client.New(m.env.ServerKey, midtrans.Sandbox)

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  booking.ID,
			GrossAmt: int64(booking.GrossAmount),
		},
		Items: &[]midtrans.ItemDetails{
			midtrans.ItemDetails{
				ID:    strconv.Itoa(int(service.ID)),
				Name:  service.Name,
				Price: int64(service.Price),
				Qty:   1,
			},
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: user.FullName,
			Email: user.Email,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		},
	}

	resp, _ := m.client.ChargeTransaction(chargeReq)

	return resp, nil
}

func (m MidTrans) CheckTransactionStatus(ctx *gin.Context) {
	bookId := ctx.Param("bookId")

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, err := m.client.CheckTransaction(bookId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get transaction status", err)
		return
	} else {
		if transactionStatusResp != nil {
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}

	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Write([]byte("ok"))
}
