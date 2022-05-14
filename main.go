package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.quick.start/database"
	"go.quick.start/register"
	"go.quick.start/routes"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	"time"
)

var validate *validator.Validate

func main() {
	godotenv.Load()
	//cont := &controller.UserController{}
	// DB connection setup
	db := database.ConnectDB()
	register.Models(db)
	//Routes
	r := routes.Setup()

	//scheduler()
	fmt.Println("Server starting at port " + os.Getenv("APP_PORT"))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
}

func scheduler() {
	/*s := gocron.NewScheduler(time.Local)
	s.Every(30).Seconds().Do(task)
	s.Every(30).Seconds().Do(vijay)

	s.Every(1).Day().At("01:14").Do(task)
	s.Every(1).Day().At("01:15").Do(task)

	s.StartAsync()*/
}

func task() {
	fmt.Println("I am runnning task.", time.Now())
}
