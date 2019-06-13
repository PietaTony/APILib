package main

import . "github.com/PietaTony/APILib/SMTP"

func main(){
	fromMail := "sjmtony@gmail.com"
	toMail := "sjmtony@gmail.com"
    subj := "This is the email subject"
    body := "This is an example body.\n With two lines."
    SMTPName := "smtp.gmail.com:465"
    SMTPMail := "sjmtony@gmail.com"
    SMTPPassword := "Sjm778887"

	SMTPSend( fromMail, toMail,
			  subj, body,
			  SMTPName,
			  SMTPMail, SMTPPassword )
}