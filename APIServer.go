package main

import (
	"github.com/PietaTony/APILib/SMTP"
	"github.com/gin-gonic/gin"
    "net/http"
)

func main(){
    APIServer("80")
}

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

func allowCrossDomain(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

    c.Next()
}

// http://pieta.ml/v1/SMTP/send?from=sjmtony@gmail.com&to=sjmtony@gmail.com&subj=tes%20t%202%202%202&body=test1&SMTPServer=smtp.gmail.com:465&SMTPMail=sjmtony@gmail.com&SMTPPassword=Sjm778887
func SMTPSend(c *gin.Context){
	from := c.Query("from")
	to := c.Query("to")
	subj := c.Query("subj")
	body := c.Query("body")
	SMTPServer := c.Query("SMTPServer")
	SMTPMail := c.Query("SMTPMail")
	SMTPPassword := c.Query("SMTPPassword")
	SMTP.Send( from, to,
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
	})
}
