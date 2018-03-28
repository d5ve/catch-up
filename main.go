package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/namsral/flag"
	"time"
)

type CatchUp struct {
	gorm.Model
	Name      string
	Details   string
	StartDate time.Time
	EndDate   time.Time
}

type Option struct {
	gorm.Model
	CatchUp CatchUp
	Date    time.Time
}

func main() {
	fmt.Println("At top of main()")

	var port int
	flag.IntVar(&port, "port", 3000, "webservice port")
	flag.Parse()
	fmt.Println("port:", port)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&CatchUp{}, &Option{})

}
