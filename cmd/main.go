package main

import (
	"github.com/crisantiz/calendar-golang/database"
	"github.com/crisantiz/calendar-golang/helpers"
)

func main() {
	helpers.LoadEnv()

	database.GetMongoConnection()
}
