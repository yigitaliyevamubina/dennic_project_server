package entity

import "time"

type DoctorWorkingHours struct {
	Id         int32
	DoctorId   string
	DayOfWeek  string
	StartTime  string
	FinishTime string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type ListDoctorWorkingHours struct {
	DoctorWhs []DoctorWorkingHours
	Count     int32
}
