package mailservice

import (
	"fmt"
	"net/smtp"

	"github.com/gondsuryaprakash/gondpariwar/logger"
	"github.com/gondsuryaprakash/gondpariwar/utils"
)

func SendMail(otp int, to ...string) error {

	funcName := "mailservice.SendMail"
	logger.I(funcName)
	// from is senders email address

	// we used environment variables to load the
	// email address and the password from the shell
	// you can also directly assign the email address
	// and the password

	from := utils.GetConfigValue("FROM_MAIL")
	password := utils.GetConfigValue("MAIL_PASSWORD")
	logger.D("From", from)
	logger.D("SendMail", password)
	// toList is list of email address that email is to be sent.
	toList := to

	// host is address of server that the
	// sender's email address belongs,
	// in this case its gmail.
	// For e.g if your are using yahoo
	// mail change the address as smtp.mail.yahoo.com
	host := "smtp.gmail.com"

	// Its the default port of smtp server
	port := "587"

	// This is the message to send in the mail
	msg := "Hello geeks!!!"

	// We can't send strings directly in mail,
	// strings need to be converted into slice bytes
	body := []byte(msg)

	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", from, password, host)

	// SendMail uses TLS connection to send the mail
	// The email is sent to all address in the toList,
	// the body should be of type bytes, not strings
	// This returns error if any occured.
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	// handling the errors
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
