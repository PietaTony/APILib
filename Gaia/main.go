package main 
import ( 
	"github.com/gin-gonic/gin"
	"net/http" 
	"fmt"
	"log"
	"context"
	"math/rand"
	"io/ioutil"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"encoding/base64"
	"time"
) 

func main() { 
	Init()
}

var (
    googleOauthConfig *oauth2.Config
)

func Init() {
	r := gin.Default()

	r.Use(allowCrossDomain)

	auth := r.Group("/auth")
	{
		auth.GET("/google/login", oauthGoogleLogin)
		auth.GET("/google/callback", oauthGoogleCallback)
	}

	api := r.Group("/api") 
	{
		api.GET("/v1/zh-TW/navbar.json", navberJson)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,  
			"error": "404, page not exists!",
		})
	})

	r.Run(":80")
}

func allowCrossDomain(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	c.Next()
}

func navberJson(c *gin.Context) {

	brand := "GAIA"
	xAxisData := []int{120, 240, rand.Intn(500), rand.Intn(500), 150, 230, 180}
	c.JSON(200, gin.H{  
		"brand": brand, 
		"xAxis_data": xAxisData,
	})
	fmt.Println("send navberJson")
}

func initOauth(){
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost/auth/google/callback",
		ClientID:     "960770367712-qji4e0i71k19adfc9vceb2cner4d3qog.apps.googleusercontent.com",
		ClientSecret: "HNd8Ise5cxCLZHQKGDc32M3z",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func oauthGoogleLogin(c *gin.Context) {
	oauthState := generateStateOauthCookie(c.Writer)

	initOauth()
	u := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(c.Writer, c.Request, u, http.StatusTemporaryRedirect)
}

func oauthGoogleCallback(c *gin.Context) {
	oauthState, _ := c.Request.Cookie("oauthstate")

	if c.Request.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(c.Request.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println( "UserInfo:\n", string(data))
	http.Redirect(c.Writer, c.Request, "https://pieta.ml", 301)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}
