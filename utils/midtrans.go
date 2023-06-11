package utils

import (
	"Capstone/models"
	"os"

	"github.com/veritrans/go-midtrans"
)

func GetPaymentURL(transaction *models.Transaction, user *models.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("Server_Key")
	midclient.ClientKey = os.Getenv("Client_Key")
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReg := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.FirstName + " " + user.LastName,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.OrderID, // masalah order ID nya gak kebaca
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReg)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}
