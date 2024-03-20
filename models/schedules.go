package models

import (
	"fmt"
	"sort"
	"time"
)

const DATETIME_FORMAT = "2006-01-02"
const TIME_FORMAT = "15:04"

type TimesScheduleRequest struct {
	StudentID    int `json:"studentId"`
	NumberOfDays int `json:"numberOfDays"`
	TypeClass    int `json:"typeClass"`
}

type TimesScheduleResponse []struct {
	ClassType struct {
		ID                      int         `json:"id"`
		CourseID                int         `json:"course_id"`
		Name                    string      `json:"name"`
		Type                    int         `json:"type"`
		LastClassOfBook         interface{} `json:"last_class_of_book"`
		Book                    int         `json:"book"`
		Next                    interface{} `json:"next"`
		Enabled                 int         `json:"enabled"`
		NumberOfStudents        int         `json:"number_of_students"`
		ComplementaryClass      int         `json:"complementary_class"`
		ComplementaryClassOrder int         `json:"complementary_class_order"`
		ScheduleByStudent       int         `json:"schedule_by_student"`
		Duration                int         `json:"duration"`
		Interval                int         `json:"interval"`
		Programmable            int         `json:"programmable"`
		CreatedAt               interface{} `json:"created_at"`
		UpdatedAt               interface{} `json:"updated_at"`
		Description             interface{} `json:"description"`
		Cor                     interface{} `json:"cor"`
		PieceOfCake             int         `json:"piece_of_cake"`
		LetsRock                int         `json:"lets_rock"`
		DoYourBest              int         `json:"do_your_best"`
		Chapter                 interface{} `json:"chapter"`
	} `json:"classType"`
	AvailableClasses []struct {
		Date                  string `json:"date"`
		ClassID               int    `json:"classId"`
		DateClassName         string `json:"dateClassName"`
		ArrayOfClassByHourDto []struct {
			Hour                  string `json:"hour"`
			ArrayOfClassByTeacher []struct {
				Teacher struct {
					ID          int         `json:"id"`
					Name        string      `json:"name"`
					Nickname    string      `json:"nickname"`
					Email       interface{} `json:"email"`
					Admission   string      `json:"admission"`
					Hours       int         `json:"hours"`
					Resignation interface{} `json:"resignation"`
					Notes       interface{} `json:"notes"`
					Enabled     int         `json:"enabled"`
					UnitID      int         `json:"unit_id"`
					UserID      interface{} `json:"user_id"`
					Photo       interface{} `json:"photo"`
					CreatedAt   string      `json:"created_at"`
					UpdatedAt   string      `json:"updated_at"`
					Link        string      `json:"link"`
					Type        string      `json:"type"`
					Pivot       struct {
						ClassID   int `json:"class_id"`
						TeacherID int `json:"teacher_id"`
					} `json:"pivot"`
				} `json:"teacher"`
				NumberOfStudents int    `json:"numberOfStudents"`
				TypeClass        string `json:"typeClass"`
				StudentIn        bool   `json:"studentIn"`
			} `json:"arrayOfClassByTeacher"`
			RetroactiveClass bool `json:"retroactiveClass"`
		} `json:"arrayOfClassByHourDto"`
	} `json:"availableClasses"`
	UnfulfilledClasses []interface{} `json:"unfulfilledClasses"`
}

type MetaData struct {
	ClassID   int
	TeacherID int
}
type Schedule struct {
	Lesson           string
	Date             string
	Time             string
	Teacher          string
	Link             string
	NumberOfStudents int
	MetaData         MetaData
}

type Schedules struct {
	Schedule []Schedule
}

func (s *Schedule) GetDateTime() string {
	return s.Date + " " + s.Time
}

func (s *Schedules) SortSchedules() {
	SortSchedules(*s)
}

func SortSchedules(s Schedules) []Schedule {
	sort.Slice(s.Schedule, func(i, j int) bool {
		return s.Schedule[i].GetDateTime() < s.Schedule[j].GetDateTime()
	})
	return s.Schedule
}

func (s *Schedules) Filter() {
	var result []Schedule
	for _, schedule := range s.Schedule {
		if schedule.roles() {
			result = append(result, schedule)
		}
	}
	s.Schedule = result
}

func (s *Schedule) roles() bool {
	datetime, err := time.Parse(DATETIME_FORMAT, s.Date)
	if err != nil {
		panic(err.Error())
	}
	hour, err := time.Parse(TIME_FORMAT, s.Time)
	if err != nil {
		panic(err.Error())
	}

	if datetime.Weekday() != 6 {
		if hour.Hour() < 9 || hour.Hour() >= 18 {
			return true
		}
		return false
	}
	return true
}

func (s *Schedules) FilterDiaSeguinte() []Schedule {
	var result []Schedule
	day := time.Now().AddDate(0, 0, 1).Day()
	for _, schedule := range s.Schedule {
		if schedule.rolesDay(day) {
			result = append(result, schedule)
		}
	}
	return result
}

func (s *Schedule) rolesDay(day int) bool {
	datetime, err := time.Parse(DATETIME_FORMAT, s.Date)
	if err != nil {
		panic(err.Error())
	}
	return datetime.Day() == day
}

func (s *Schedules) SprintSchedules() string {
	var result string
	for _, schedule := range s.Schedule {
		result += schedule.SprintSchedule()
	}
	return result
}

func (s *Schedule) SprintSchedule() string {
	return fmt.Sprintln("-----------------------------------------------") +
		fmt.Sprintln("Aula:", s.Lesson) +
		fmt.Sprintln("Dia:", s.Date, s.Time) +
		fmt.Sprintln("Professor:", s.Teacher) +
		fmt.Sprintln("Link:", s.Link) +
		fmt.Sprintln("Quantidade de alunos matriculados:", s.NumberOfStudents) +
		fmt.Sprintln("")
}
