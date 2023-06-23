package main

import (
	"fmt"
	"database/sql"
	util "github.com/naviscom/simplebank/util"
	"log"

	_ "github.com/lib/pq"

)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Println("cannot load config")
	}

	if config.Environment == "development" {
		fmt.Println("Environment = development")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	//connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Println("cannot connect to db")
	}

	fmt.Println("connected to db")
	conn.Close()
}
