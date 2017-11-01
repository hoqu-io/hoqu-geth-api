package main

import (
	"os"
	"hoqu-geth-api/http"
	"hoqu-geth-api/sdk/app"
	httpSdk "hoqu-geth-api/sdk/http"
	"hoqu-geth-api/geth"
    "github.com/sirupsen/logrus"
)

func init() {
	app.InitConfig()
	app.InitLogger()
	if err := geth.InitGeth(); err != nil {
	    logrus.Fatal(err)
    }
}

func main() {
	interrupt := make(chan os.Signal, 1)

	go http.Run()

	httpSdk.StartTicking(interrupt)
}