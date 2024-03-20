package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"os"
	"projectIntern/internal/entity"
	"strconv"
)

type MidTransItf interface {
	CreateTransactionService(service entity.Service, paymentMethod string, bookingId string, grossAmt int64, user entity.User) (*coreapi.ChargeResponse, error)
	CreateTransactionDoctor(doctor entity.Doctor, paymentMethod string, bookingId string, grossAmt int64, user entity.User) (*coreapi.ChargeResponse, error)
	Notification(bookingId string) (*coreapi.TransactionStatusResponse, error)
}

type MidTrans struct {
	client coreapi.Client
}

func NewMidtrans(client coreapi.Client) MidTransItf {
	return &MidTrans{client: client}
}

func (m MidTrans) CreateTransactionService(service entity.Service, paymentMethod string, bookingId string, grossAmt int64, user entity.User) (*coreapi.ChargeResponse, error) {
	m.client.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  bookingId,
			GrossAmt: grossAmt,
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
	}

	if paymentMethod == "bca" {
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		}
	} else if paymentMethod == "bni" {
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBni,
		}
	}

	resp, _ := m.client.ChargeTransaction(chargeReq)

	return resp, nil
}

func (m MidTrans) CreateTransactionDoctor(doctor entity.Doctor, paymentMethod string, bookingId string, grossAmt int64, user entity.User) (*coreapi.ChargeResponse, error) {
	m.client.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  bookingId,
			GrossAmt: grossAmt,
		},
		Items: &[]midtrans.ItemDetails{
			midtrans.ItemDetails{
				ID:    doctor.ID.String(),
				Name:  doctor.Name,
				Price: int64(doctor.Price),
				Qty:   1,
			},
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: user.FullName,
			Email: user.Email,
		},
	}

	if paymentMethod == "bca" {
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		}
	} else if paymentMethod == "bni" {
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBni,
		}
	}

	resp, _ := m.client.ChargeTransaction(chargeReq)

	return resp, nil
}

func (m MidTrans) Notification(bookingId string) (*coreapi.TransactionStatusResponse, error) {
	m.client.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	check, err := m.client.CheckTransaction(bookingId)
	if err != nil {
		return nil, err
	}

	return check, nil
}
