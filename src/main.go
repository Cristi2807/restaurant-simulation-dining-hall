package main

import (
	"bytes"
	"github.com/beego/beego/v2/server/web"
	"log"
	"net/http"
)

func main() {
	web.Router("/distribution", &MainController{})
	web.Run()
}

type MainController struct {
	web.Controller
}

func (this *MainController) Post() {
	log.Println("received in dining hall")

	jsonBody := []byte(``)
	bodyReader := bytes.NewReader(jsonBody)

	_, _ = http.Post("http://restaurant-kitchen:8050/order", "application/json", bodyReader)
}
