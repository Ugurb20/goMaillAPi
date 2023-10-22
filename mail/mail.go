package mail

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

///ldvv moar jqnu yhkd

type Contact struct {
	Email   string `json:"email"`
	CompanyName string `json:"companyName"`
	FirstName   string `json:"firstName"`
}


func SendMail(contact Contact) (string,error){

	msg := gomail.NewMessage()
    msg.SetHeader("From", "ugurbukcuoglu@gmail.com")
    msg.SetHeader("To", "ugurbukcuoglu@gmail.com")
    msg.SetHeader("Subject", "New Lead")
    msg.SetBody("text/html", fmt.Sprintf("<b><p>Email: %s</p><p>Firstname: %s</p><p>Company Name: %s</p></b>",contact.Email, contact.FirstName, contact.CompanyName))

    n := gomail.NewDialer("smtp.gmail.com", 587, "ugurbukcuoglu@gmail.com", "zdtklxazocobooso")

    // Send the email
    if err := n.DialAndSend(msg); err != nil {
        panic(err)
    }
	
	return "Email Sent Successfully!", nil
}