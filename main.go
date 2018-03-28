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
	Options   []Option
}

type Option struct {
	gorm.Model
	CatchUpID int
	Date      time.Time
	Votes     []Vote
}

type Vote struct {
	gorm.Model
	YNM    string `gorm:"size:1"`
	person string
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

	db.AutoMigrate(&CatchUp{}, &Option{}, &Vote{})

	c1 := CatchUp{
		Name:      "First catchup",
		Details:   "First attempt at a catchup record",
		StartDate: time.Date(2018, time.March, 10, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2018, time.March, 17, 0, 0, 0, 0, time.UTC),
	}

	db.Create(&c1)
}
