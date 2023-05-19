package usecase

import (
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"Capstone/utils"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/mailjet/mailjet-apiv3-go"

	"golang.org/x/crypto/bcrypt"
)

var generatedOTP string

func LoginUser(req *payload.LoginUserRequest) (res payload.LoginUserResponse, err error) {

	user, err := database.GetuserByEmail(req.Email)
	if err != nil {
		return res, echo.NewHTTPError(http.StatusBadRequest, "Email not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("Wrong Password")
	}

	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		return res, echo.NewHTTPError(http.StatusBadRequest, "Failed to generate token")
	}

	user.Token = token

	res = payload.LoginUserResponse{
		Email: user.Email,
		Token: user.Token,
	}

	return
}

func SendOTPByEmail(emailAddress, otp string) error {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: os.Getenv("MJ_FROM_EMAIL"),
				Name:  "INVENTRON support",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: emailAddress,
					Name:  "",
				},
			},
			Subject:  "OTP for reset password",
			TextPart: "Dear our costumer, dont share below OTP code if you have",
			HTMLPart: fmt.Sprintf("<h3>Kode OTP kamu adalah <span>%s</span>, berlaku selama 5 menit.</h3>", otp),
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}

	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
		fmt.Println(res)
	}
	return nil
}
func GenerateOTPEndpoint(req *payload.ForgotPasswordRequest) error {
	user, err := database.GetuserByEmail(req.Email)
	if err != nil {
		return errors.New("Email not registered")
	}
	generatedOTP = utils.GenerateOTP()

	err = SendOTPByEmail(user.Email, generatedOTP)
	if err != nil {
		return errors.New("Failed to send OTP")
	}
	return nil
}

func VerifyOTP(req *payload.VerifyngOtp) error {
	if req.Otp != generatedOTP {
		return errors.New("OTP verification failed.")
	}

	return errors.New("OTP verification successful!")
}
