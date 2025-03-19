package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application , error){
	logger := log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I am alive!")
}