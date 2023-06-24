package main

import (
	"database/sql"
	"log"

	// "net/http"
	"os"
	// "time"

	util "github.com/naviscom/simplebank/util"

	_ "github.com/lib/pq"
)

// type application struct {
// 	logger *log.Logger
// //	models models.Models
// }

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	config, err := util.LoadConfig(".")
	if err != nil {
		logger.Println("cannot load config")
	}
	if config.Environment == "development" {
		logger.Println("Environment = development")
	}
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		logger.Println("cannot connect to db")
	}
	logger.Println("connected to db")
	defer db.Close()

	// app := &application{
	// 	logger: logger,
	// 	models: models.NewModels(db),
	// }
	// srv := &http.Server{
	// 	Addr:         fmt.Sprintf(":%d", config.HTTPServerPort),
	// 	Handler:      app.routes(),
	// 	IdleTimeout:  time.Minute,
	// 	ReadTimeout:  10 * time.Second,
	// 	WriteTimeout: 30 * time.Second,
	// }
	// logger.Println("Starting api server on port", config.HTTPServerAddress)
	// err = srv.ListenAndServe()
	// if err != nil {
	// 	log.Println(err)
	// }
}
