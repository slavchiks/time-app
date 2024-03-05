package endpoint

import (
	"fmt"
	"net/http"
)

type Service interface{
	DaysLeft() int 
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint{
	return &Endpoint{
		s:s,
	}
}


func (e *Endpoint) Status(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		d:= e.s.DaysLeft()
		fmt.Fprintf(w,"Days left : %d", d)
	}
}