package main

import (
	"context"
	"fmt"

	"github.com/hytech-racing/Mock-Cloud-Server/app"
)

func main() {
	a := app.New()

	err := a.Start(context.TODO())

	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
