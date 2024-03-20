package integrations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MauricioUhlig/times-classes-go/models"
)

func DoRequest() []byte {
	var requestBody = models.TimesScheduleRequest{StudentID: 0 /*id*/, NumberOfDays: 14, TypeClass: 1 /*online*/}

	var posturl = "<URL>"
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Authorization", "Bearer <TOKEN>")

	var c http.Client
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return response
}
func Unmarshal(response []byte) models.TimesScheduleResponse {
	var timesScheduleResponse models.TimesScheduleResponse
	err := json.Unmarshal(response, &timesScheduleResponse)
	if err != nil {
		fmt.Println(err)
	}
	return timesScheduleResponse
}

func Print(available *models.Schedules, scheduled *models.Schedules) {
	available.SortSchedules()
	fmt.Println("Disponíveis")
	fmt.Println(available.SprintSchedules())
	fmt.Println("")
	fmt.Println("Agendados")
	scheduled.SortSchedules()
	fmt.Println(scheduled.SprintSchedules())
}
func Sprint(available *models.Schedules, scheduled *models.Schedules) string {
	available.SortSchedules()
	scheduled.SortSchedules()
	return fmt.Sprintln("Disponíveis") +
		available.SprintSchedules() +
		fmt.Sprintln("") +
		fmt.Sprintln("Agendados") +
		scheduled.SprintSchedules()
}
func FindClass(schedules *models.TimesScheduleResponse) (models.Schedules, models.Schedules) {
	var availables models.Schedules
	var appointments models.Schedules
	for _, schedule := range *schedules {
		for _, availableClass := range schedule.AvailableClasses {
			for _, classByHourDto := range availableClass.ArrayOfClassByHourDto {
				for _, classByTeacher := range classByHourDto.ArrayOfClassByTeacher {
					if classByTeacher.StudentIn {
						s := models.Schedule{
							Lesson:           schedule.ClassType.Name + " " + availableClass.DateClassName,
							Date:             availableClass.Date,
							Time:             classByHourDto.Hour,
							Teacher:          classByTeacher.Teacher.Name,
							Link:             classByTeacher.Teacher.Link,
							NumberOfStudents: classByTeacher.NumberOfStudents,
							MetaData: models.MetaData{
								ClassID:   schedule.ClassType.ID,
								TeacherID: classByTeacher.Teacher.ID,
							},
						}
						appointments.Schedule = append(appointments.Schedule, s)
					} else if classByTeacher.TypeClass == "online" {
						availables.Schedule = append(availables.Schedule, models.Schedule{
							Lesson:           schedule.ClassType.Name + " " + availableClass.DateClassName,
							Date:             availableClass.Date,
							Time:             classByHourDto.Hour,
							Teacher:          classByTeacher.Teacher.Name,
							Link:             classByTeacher.Teacher.Link,
							NumberOfStudents: classByTeacher.NumberOfStudents,
							MetaData: models.MetaData{
								ClassID:   schedule.ClassType.ID,
								TeacherID: classByTeacher.Teacher.ID,
							},
						},
						)
					}
				}
			}

		}
	}
	return availables, appointments
}
