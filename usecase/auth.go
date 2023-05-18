package usecase

import (
	"Capstone/constants"
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"Capstone/utils"
	"errors"
	"fmt"
	"log"

	// "os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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
	err := godotenv.Load("utils.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	from := mail.NewEmail(constants.SEND_FROM_NAME, constants.SEND_FROM_ADDRESS)
	subject := "INVENTORN OTP RESET PASSWORD"
	to := mail.NewEmail("Recipient Name", emailAddress)
	plainTextContent := fmt.Sprintf("dont you dare to give this code to other people, Your OTP is: %s", otp)
	htmlContent := fmt.Sprintf("<strong>Your OTP is: %s </strong>", plainTextContent)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(constants.SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		fmt.Println("Unable to send your email")
		log.Fatal(err)
	}

	statusCode := response.StatusCode
	if statusCode == 200 || statusCode == 201 || statusCode == 202 {
		fmt.Println("Email sent!")
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
