package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/cisdi-dev/soal-techlead-be-1/modules/hospitals"
	"github.com/cisdi-dev/soal-techlead-be-1/modules/patients"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@127.0.0.1:3306/db_rumkit")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	mux := http.NewServeMux()

	hospitalsRepo := hospitals.NewRepository(db)
	hospitals.NewDelivery(mux, "/hospitals", hospitals.NewUsecase(hospitalsRepo))

	patientsRepo := patients.NewRepository(db)
	patients.NewDelivery(mux, "/patients", patients.NewUsecase(patientsRepo, hospitalsRepo))

	http.ListenAndServe(":5000", mux)
}
