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

	// Mon Jan 2 15:04:05 -0700 MST 2006
	d1 := "2018-03-10"
	start_date, err := time.Parse("2006-01-02", d1)
	if err != nil {
		panic(fmt.Sprint("failed to parse ", d1))
	}

	d2 := "2018-03-17"
	end_date, err := time.Parse("2006-01-02", d2)
	if err != nil {
		panic(fmt.Sprint("failed to parse ", d2))
	}

	options := getOptions(start_date, end_date)

	c1 := CatchUp{
		Name:      "First catchup",
		Details:   "First attempt at a catchup record",
		StartDate: start_date,
		EndDate:   end_date,
		Options:   options,
	}

	db.Create(&c1)
}

func getOptions(StartDate time.Time, EndDate time.Time) []Option {
	if StartDate.After(EndDate) {
		StartDate, EndDate = EndDate, StartDate
	}
	var options []Option
	for day := StartDate; day.Before(EndDate) || day.Equal(EndDate); day = day.AddDate(0, 0, 1) {
		o := Option{Date: day}
		options = append(options, o)
	}
	return options
}
