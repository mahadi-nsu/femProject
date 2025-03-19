package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/mahadi-nsu/femProject/internal/app"
	"github.com/mahadi-nsu/femProject/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
    
	app.Logger.Printf("We are running out app at port %d\n", port)
	
	r := routes.SetupRoutes(app)


	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

