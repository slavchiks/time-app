package endpoint

import (
	"fmt"
	"net/http"
	"time"
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
		/*http.ServeFile(w, r, "form.html")
		if r.Method == "POST" {
            dateStr := r.FormValue("date")
            layout := "2006-01-02"
			date, err := time.Parse(layout, dateStr)
				if err != nil {
					http.Error(w, "Ошибка при парсинге даты", http.StatusBadRequest)
					return
				}
				if date.After(time.Now()) {
					fmt.Fprintf(w, "Введенная дата %s находится в будущем", dateStr)
				} else {
					fmt.Fprintf(w, "Введенная дата %s находится в прошлом или сегодня", dateStr)
				}
		}*/
		d:= e.s.DaysLeft()
		fmt.Fprintf(w,"Days left : %d", d)
	}
}