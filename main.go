package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/namsral/flag"
	"log"
	//"os"
)

func main() {
	fmt.Println("At top of main()")

	var port int
	flag.IntVar(&port, "port", 3000, "webservice port")
	flag.Parse()
	fmt.Println("port:", port)

	db, err := sql.Open("sqlite3", "./catch-up.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
