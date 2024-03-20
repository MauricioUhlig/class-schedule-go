package main

import (
	"fmt"

	"github.com/MauricioUhlig/times-classes-go/integrations"
	"github.com/MauricioUhlig/times-classes-go/models"
)

func main() {
	response := integrations.DoRequest()
	timesResponse := integrations.Unmarshal(response)
	available, scheduled := integrations.FindClass(&timesResponse)
	available.Filter()
	availableTomorow := models.Schedules{Schedule: available.FilterDiaSeguinte()}
	scheduledTomorow := models.Schedules{Schedule: scheduled.FilterDiaSeguinte()}
	fmt.Println(integrations.Sprint(&availableTomorow, &scheduledTomorow))
}
