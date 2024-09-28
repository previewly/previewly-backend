package main

import (
	"wsw/backend/app"
	"wsw/backend/lib/utils"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		utils.F("Coulnd not create application: %v", err)
	}
	application.Start()
	defer application.Closer()()
}
