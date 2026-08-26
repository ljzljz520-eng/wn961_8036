package main

import (
	"context"
	"frontend_go/internal/api"
	"frontend_go/internal/service"
	"frontend_go/internal/store"
	"log"
	"net/http"
)

func main() {
	s, e := store.Open("training.db")
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	svc := service.New(s)
	h := api.New(svc)
	_ = context.Background()
	log.Fatal(http.ListenAndServe(":8080", h))
}
