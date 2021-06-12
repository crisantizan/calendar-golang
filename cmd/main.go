package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/crisantiz/calendar-golang/helpers"
	"github.com/crisantiz/calendar-golang/repositories"
)

func main() {
	helpers.LoadEnv()

	calendarRepo := repositories.New()

	calendar, err := calendarRepo.GetOne("60a18a867bb85dcedaeff729")

	if err != nil {
		log.Fatal(err)
	}

	e, err := json.Marshal(&calendar)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(e))
}
