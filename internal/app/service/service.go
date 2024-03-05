package service

import (
	"net/http"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}


func (s *Service) DaysLeft() int {
	data := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(data)
	return int(dur.Hours())/24
}
func (s *Service) DiffDays(data time.Time) int64 {
	dur := time.Until(data)
	return int64(dur.Hours())/24
}
func (s *Service)ParsDate(r *http.Request)(time.Time, error){
	dateStr := r.FormValue("date")
	layout := "2006-01-02"
	date, err := time.Parse(layout, dateStr)
	return date, err
}