package main

import (
	"github.com/PietaTony/APILib/SMTP"
	"github.com/gin-gonic/gin"
    "net/http"
)

func main(){
    APIServer("80")
}

/*
API Server 運行以及呼叫分類
*/
func APIServer(port string){
    engine := gin.Default()

    v1 := engine.Group("/v1", allowCrossDomain)
    {
        SMTP := v1.Group("/SMTP")
        {
        	SMTP.GET("/send", SMTPSend)
        }
    }

    engine.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "status": 404,  
            "error": "404, page not exists!",
        })
    })

    engine.Run(":" + port)
}

/*
CORS請求
*/
func allowCrossDomain(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

    c.Next()
}

/*
將POST到的資訊寄送出
"from":寄信端(Email string)
"to": 收信端(Email string)
"subj": 主題(string)
"body": 內容(string)
"SMTPServer": SMTP的伺服器
"SMTPMail": SMTP帳號(Email string)
"SMTPPassword": SMTP密碼(Email password string)
*/
func SMTPSend(c *gin.Context){
	from := c.Query("from")
	to := c.Query("to")
	subj := c.Query("subj")
	body := c.Query("body")
	SMTPServer := c.Query("SMTPServer")
	SMTPMail := c.Query("SMTPMail")
	SMTPPassword := c.Query("SMTPPassword")
	success := SMTP.Send( from, to,
			  subj, body,
			  SMTPServer,
			  SMTPMail, SMTPPassword )
	c.JSON(200, gin.H{
		"from": from,
		"to": to,
		"subj": subj,
		"body": body,
		"SMTPServer": SMTPServer,
		"SMTPMail": SMTPMail,
		"SMTPPassword": SMTPPassword,
		"success": success,
	})
}
