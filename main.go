package main

import (
	"fmt"
	"github.com/thayen/sample-rest/domain/service"
	http2 "github.com/thayen/sample-rest/interface/http"
	"github.com/thayen/sample-rest/interface/persistence"
	"github.com/thayen/sample-rest/usecase"
	"log"
	"net/http"
	"os"
)

func main() {

	log.SetFlags(log.Lshortfile)
	port := getOrDefault("PORT", "8080")

	repository, err := persistence.NewContactSqlRepository(getOrDefault("POSTGRES_URL", "postgres://postgres:postgres@127.0.0.1:5432/?sslmode=disable"))
	if err != nil {
		log.Fatal("cannot create repository ", err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), http2.NewHttpContact(usecase.NewContactUsecase(service.NewContactService(repository), repository))); err != nil {
		log.Fatal(err)
	}
}

func getOrDefault(env string, defaultvalue string) string {
	if p := os.Getenv(env); p != "" {
		return p
	}
	return defaultvalue
}
