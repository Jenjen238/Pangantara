package payment

import (
	"fmt"
	"sppg-backend/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

var snapClient snap.Client
var coreClient coreapi.Client

func InitMidtrans() {
	env := midtrans.Sandbox
	if config.AppConfig.MidtransEnv == "production" {
		env = midtrans.Production
	}

	snapClient.New(config.AppConfig.MidtransServerKey, env)
	coreClient.New(config.AppConfig.MidtransServerKey, env)
}

type PaymentRequest struct {
	OrderID       string
	Amount        int64
	CustomerName  string
	CustomerEmail string
	Items         []PaymentItem
}

type PaymentItem struct {
	ID       string
	Name     string
	Price    int64
	Quantity int
}

type PaymentResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

func CreatePayment(req PaymentRequest) (*PaymentResponse, error) {
	var itemDetails []midtrans.ItemDetails
	for _, item := range req.Items {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
			Qty:   int32(item.Quantity),
		})
	}

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderID,
			GrossAmt: req.Amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: req.CustomerName,
			Email: req.CustomerEmail,
		},
		Items: &itemDetails,
		Expiry: &snap.ExpiryDetails{
			Unit:     "hour",
			Duration: 24,
		},
	}

	snapResp, err := snapClient.CreateTransaction(snapReq)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat payment: %v", err)
	}

	return &PaymentResponse{
		Token:       snapResp.Token,
		RedirectURL: snapResp.RedirectURL,
	}, nil
}

func VerifyPayment(orderID string) (string, error) {
	transactionStatus, err := coreClient.CheckTransaction(orderID)
	if err != nil {
		return "", fmt.Errorf("gagal cek status payment: %v", err)
	}

	return transactionStatus.TransactionStatus, nil
}