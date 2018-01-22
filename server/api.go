package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"crypto/rand"
	"encoding/base64"

	"os"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// API is a defined as struct bundle
// for api. Feel free to organize
// your app as you wish.
type API struct{}

// Bind attaches api routes
func (api *API) Bind(group *echo.Group) {
	group.GET("/v1/conf", api.ConfHandler)
	group.GET("/v1/auth", api.GoogleAuth)
}

// ConfHandler handle the app config, for example
func (api *API) ConfHandler(c echo.Context) error {
	app := c.Get("app").(*App)
	return c.JSON(200, app.Conf.Root)
}

// I try to API
func (api *API) GoogleAuth(c echo.Context) error {
	//app := c.Get("app").(*App)

	return c.String(200, LoadCreds())
}

// Credentials which stores google ids.
type Credentials struct {
    Cid string `json:"cid"`
    Csecret string `json:"csecret"`
}

func LoadCreds() string{
    var c Credentials
    file, err := ioutil.ReadFile("./creds.json")
    if err != nil {
        fmt.Printf("File error: %v\n", err)
        os.Exit(1)
    }
    json.Unmarshal(file, &c)
	conf := &oauth2.Config{
	  ClientID:     c.Cid,
	  ClientSecret: c.Csecret,
	  RedirectURL:  "http://127.0.0.1:5001/authorized",
	  Scopes: []string{
	    "https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/spreadsheets",	 // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
	  },
	  Endpoint: google.Endpoint,
	}

	token := randToken()

	result := getLoginURL(token, conf)
	return result
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func getLoginURL(state string, conf *oauth2.Config) string {
    // State can be some kind of random generated hash string.
    // See relevant RFC: http://tools.ietf.org/html/rfc6749#section-10.12
    return conf.AuthCodeURL(state)
}
