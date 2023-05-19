package usecase

import (
	"Capstone/constants"
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"Capstone/utils"
	"errors"
	"fmt"
	// "log"
	// "os"
	// "os"
	"github.com/mailjet/mailjet-apiv3-go/v4"

	// "github.com/joho/godotenv"
	"github.com/labstack/echo"
	// "github.com/sendgrid/sendgrid-go"
	// "github.com/sendgrid/sendgrid-go/helpers/mail"
	"golang.org/x/crypto/bcrypt"
)

var generatedOTP string

func LoginUser(req *payload.LoginUserRequest) (res payload.LoginUserResponse, err error) {

	user, err := database.GetuserByEmail(req.Email)
	if err != nil {
		echo.NewHTTPError(400, "Email not registered")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("Wrong Password")
	}

	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		echo.NewHTTPError(400, "Failed to generate token")
		return
	}

	user.Token = token

	res = payload.LoginUserResponse{
		Email: user.Email,
		Token: user.Token,
	}

	return
}

func SendOTPByEmail(emailAddress, otp string) error {
	mailjetClient := mailjet.NewMailjetClient(constants.MJ_APIKEY_PUBLIC, constants.MJ_APIKEY_PRIVATE)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: constants.MJ_FROM_EMAIL,
				Name:  "INVENTRON support",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: emailAddress,
					Name:  "",
				},
			},
			Subject:  "OTP for reset password",
			TextPart: "Dear our costumer, dont share below OTP code if you have ",
			HTMLPart: "<h3>Your otp code otp : </h3> " + otp ,
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
