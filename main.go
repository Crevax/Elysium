package main

import (
	"log"
	"os"

	"github.com/labstack/echo"

	"fmt"

	"cjdavis.me/elysium/controllers"
)

func main() {
	router := echo.New()

	controllers.Init(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Elysium listening on port %s", port)
	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", port)))
}
