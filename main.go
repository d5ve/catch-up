package main

//go:generate sqlboiler mysql --wipe

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/namsral/flag"
	"gopkg.in/inconshreveable/log15.v2"
)

var log = log15.New()

func main() {
	log.Info("At top of main()")

	var port int
	flag.IntVar(&port, "port", 3000, "webservice port")
	flag.Parse()
	log.Info("Got port", "port", port)

	db, err := sql.Open("mysql", "catch_up:catch_up_password@tcp(0.0.0.0:3306)/catch_up?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	log.Info("Got DB handle", "db", db)

	routes := setupRoutes(db)
	log.Info("Got routes", "routes", routes)

	routes.Run(fmt.Sprintf(":%d", port))
}
