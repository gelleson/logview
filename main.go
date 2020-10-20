package main

import (
	"fmt"
	"github.com/gelleson/logview/pkg/application"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	js := mewn.String("./frontend/dist/my-app/main.js")
	css := mewn.String("./frontend/dist/my-app/styles.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "logview",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	appInstance := application.New(application.Option{})

	if err := appInstance.Build(); err != nil {
		fmt.Println(err)
	}

	service := appInstance.Service()

	app.Bind(appInstance)
	app.Bind(service.LogService)
	app.Bind(service.UploadService)
	app.Run()
}
