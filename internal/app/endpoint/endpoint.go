package endpoint

import (
	"fmt"
	"net/http"
	"time"
)

type Service interface {
	DaysLeft() int
	DiffDays(time.Time) int
	ParsDate(*http.Request) (time.Time, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "cmd\\time-task\\form.html")
		d := e.s.DaysLeft()
		fmt.Fprintf(w, "Days left : %d", d+1)

	} else if r.Method == http.MethodPost {
		date, err := e.s.ParsDate(r)

		if err != nil {
			http.Error(w, "Date parsing error!", http.StatusBadRequest)
			return
		}

		res := e.s.DiffDays(date)
		fmt.Fprintf(w, "Days left : %d", res+1)
	}
}
