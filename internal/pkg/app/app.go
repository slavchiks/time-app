package app

import (
	"db/time/internal/app/endpoint"
	"db/time/internal/app/mw"
	"db/time/internal/app/service"
	"log"
	"net/http"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
}

func New()(*App, error){
	a:= &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)
	http.HandleFunc("/time", mw.Role_chek(a.e.Status))
	
	return a, nil
}

func (a *App) Run() error{

	err := http.ListenAndServe("localhost:8080",nil)
	if err!=nil{
		log.Fatal("Server error.")
	}
	return err
}