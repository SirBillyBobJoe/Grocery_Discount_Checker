package service

import (
	"fmt"
	"net/smtp"
	"os"

	log "github.com/sirupsen/logrus"
)

var semaphore = make(chan struct{}, 10)

func sendEmail(itemName string, originalPrice float32, newPrice float32, email string) {
	discount := originalPrice - newPrice
	discountPercent := (discount / originalPrice) * 100

	msg := "Subject: Woolworths Price Alert!\r\n\r\n" +
		"Good news! A product you're tracking at Woolworths is now cheaper.\r\n\r\n" +
		"Item: " + itemName + "\r\n" +
		"Original Price: $" + fmt.Sprintf("%.2f", originalPrice) + "\r\n" +
		"Current Price:  $" + fmt.Sprintf("%.2f", newPrice) + "\r\n" +
		"You save:       $" + fmt.Sprintf("%.2f", discount) +
		" (" + fmt.Sprintf("%.1f", discountPercent) + "%% off)\r\n\r\n"

	auth := smtp.PlainAuth(
		"", os.Getenv("SENDER_EMAIL"), os.Getenv("APP_PASSWORD"), "smtp.gmail.com",
	)

	fmt.Println("SENDING EMAIL")

	if err := smtp.SendMail(
		"smtp.gmail.com:587", auth, os.Getenv("SENDER_EMAIL"), []string{email}, []byte(msg),
	); err != nil {
		log.Error("Error sending email", err)
	}

}

func SendAsyncEmail(itemName string, originalPrice float32, newPrice float32, email string) {
	semaphore <- struct{}{}

	defer func() {
		<-semaphore
	}()

	sendEmail(itemName, originalPrice, newPrice, email)
}
