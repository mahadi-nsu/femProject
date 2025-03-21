package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mahadi-nsu/femProject/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	// Store will be here

	// Initialize handlers
	workoutHandler := api.NewWorkoutHandler()
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I am alive!")
}
