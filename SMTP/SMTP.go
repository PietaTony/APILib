package SMTP

import (
    "fmt"
	"log"
    "net"
    "net/mail"
	"net/smtp"
    "crypto/tls"
)


/*
"from":寄信端(Email string)
"to": 收信端(Email string)
"subj": 主題(string)
"body": 內容(string)
"SMTPServer": SMTP的伺服器
"SMTPMail": SMTP帳號(Email string)
"SMTPPassword": SMTP密碼(Email password string)
*/
func Send( fromMail string,
           toMail string,
           subj string,
           body string,
           SMTPServer string,
           SMTPMail string,
           SMTPPassword string ) (bool) {

    from := mail.Address{"", fromMail}
    to   := mail.Address{"", toMail}
    //subj := "This is the email subject"
    //body := "This is an example body.\n With two lines."

    // Setup headers
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subj

    // Setup message
    message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    // Connect to the SMTP Server
    //SMTPServer := "smtp.gmail.com:465"

    host, _, _ := net.SplitHostPort(SMTPServer)

    auth := smtp.PlainAuth("", SMTPMail, SMTPPassword, host)

    // TLS config
    tlsconfig := &tls.Config {
        InsecureSkipVerify: true,
        ServerName: host,
    }

    // Here is the key, you need to call tls.Dial instead of smtp.Dial
    // for smtp servers running on 465 that require an ssl connection
    // from the very beginning (no starttls)
    conn, err := tls.Dial("tcp", SMTPServer, tlsconfig)
    if err != nil {
        log.Panic(err)
        return false
    }

    c, err := smtp.NewClient(conn, host)
    if err != nil {
        log.Panic(err)
        return false
    }

    // Auth
    if err = c.Auth(auth); err != nil {
        log.Panic(err)
        return false
    }

    // To && From
    if err = c.Mail(from.Address); err != nil {
        log.Panic(err)
        return false
    }

    if err = c.Rcpt(to.Address); err != nil {
        log.Panic(err)
        return false
    }

    // Data
    w, err := c.Data()
    if err != nil {
        log.Panic(err)
        return false
    }

    _, err = w.Write([]byte(message))
    if err != nil {
        log.Panic(err)
        return false
    }

    err = w.Close()
    if err != nil {
        log.Panic(err)
        return false
    }

    c.Quit()
    fmt.Println("SMTP Send Success")
    return true
}