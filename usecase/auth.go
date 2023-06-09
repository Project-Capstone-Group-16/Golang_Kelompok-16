package usecase

import (
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"Capstone/utils"
	"errors"
	"fmt"
	"os"

	"github.com/mailjet/mailjet-apiv3-go"

	"golang.org/x/crypto/bcrypt"
)

// Logic Login User
func LoginUser(req *payload.LoginUserRequest) (res payload.LoginUserResponse, err error) {

	user, err := database.GetUserByEmail(req.Email)
	if err != nil {
		return res, errors.New("Email Not Registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("Wrong Password")
	}

	token, err := middleware.CreateToken(int(user.ID), user.Role)
	if err != nil {
		return res, errors.New("Failed To Create Token")
	}

	user.Token = token

	res = payload.LoginUserResponse{
		Email: user.Email,
		Token: user.Token,
	}

	return
}

// Logic Login Admin
func LoginAdmin(req *payload.LoginAdminRequest) (res payload.LoginAdminResponse, err error) {

	user, err := database.GetUserByEmail(req.Email)
	if err != nil {
		return res, errors.New("Email not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("Wrong Password")
	}

	token, err := middleware.CreateToken(int(user.ID), user.Role)
	if err != nil {
		return res, errors.New("Failed To Create Token")
	}

	user.Token = token

	res = payload.LoginAdminResponse{
		Email: user.Email,
		Token: user.Token,
	}

	return
}

// Logic OTP email
func SendOTPByEmail(emailAddress, otp string) error {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: os.Getenv("MJ_FROM_EMAIL"),
				Name:  "INVENTRON-no-reply",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: emailAddress,
					Name:  "",
				},
			},
			Subject:  "OTP for reset password",
			TextPart: "Dear our costumer, dont share below OTP code if you have ",
			HTMLPart: fmt.Sprintf("<html><head><style>.container{width:600px;margin:0 auto;border:1px solid #e0e0e0;border-radius:4px;padding:40px;}h1{color:#1652F9;font-weight:bold;font-family:'Poppins',sans-serif;text-align:left;}.underline{text-decoration:underline;}.footer{text-align:center;font-size:12px;color:#888888;}.line{border-top: 5px solid #e0e0e0;margin-top: 20px;margin-bottom: 20px;}body{font-family: 'Poppins', sans-serif;}</style></head><body><div class=\"container\"><h1>Inventron</h1><div class=\"line\"></div><p>We have received a request to reset the password for your account. To proceed with the password reset, please use the One-Time Password (OTP) provided below:</p><br><h1 style=\"text-align:center;\"><style=><span class=\"underline;\"><strong>%s</h1></strong></span></h2><br><p>Please note this OTP is valid for 5 minutes only. If you did not initiate this password reset request, please disregard this email.</p><p>If you are the one requesting the password reset, please enter the OTP on the password reset page to complete the process. Ensure that you keep this OTP confidential and do not share it with anyone.</p><br><p>Thank you for choosing our service.</p><br><p class=\"footer\">This email was sent by Inventron.<br>Jakarta, Indonesia<br>©2023 Capstone16Group, ALTA. | Privacy Policy</p></div></body></html>", otp),
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

// Logic Generate  OTP
func GenerateOTPEndpoint(req *payload.ForgotPasswordRequest) (res payload.GenerateOTPResponse, err error) {
	user, err := database.GetUserByEmail(req.Email)
	if err != nil {
		return res, errors.New("Email not registered")
	}

	user.OTP = utils.GenerateOTP()

	err = SendOTPByEmail(user.Email, user.OTP)
	if err != nil {
		return res, errors.New("Failed to send OTP")
	}

	err = database.UpdateUser(&user)
	if err != nil {
		return res, errors.New("Failed to update user")
	}

	res = payload.GenerateOTPResponse{
		Email: user.Email,
	}

	return res, nil
}

// Logic Verify OTP
func VerifyOTP(req *payload.VerifyngOtpRequest) (res payload.LoginUserResponse, err error) {
	user, err := database.GetUserByEmail(req.Email)
	if err != nil {
		return res, errors.New("Failed to get user")
	}

	if req.Otp != user.OTP {
		return res, errors.New("OTP verification failed.")
	}

	token, err := middleware.CreateToken(int(user.ID), user.Role)
	if err != nil {
		return res, errors.New("Failed To Create Token")
	}

	res = payload.LoginUserResponse{
		Email: user.Email,
		Token: token,
	}

	return res, nil
}
