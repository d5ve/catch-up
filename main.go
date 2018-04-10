package main

//go:generate sqlboiler mysql --wipe

import (
	"database/sql"
	"fmt"
	"github.com/d5ve/catch-up/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/namsral/flag"
	"gopkg.in/inconshreveable/log15.v2"
	"time"
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

	catch_ups, err := models.CatchUps(db).All()
	log.Info("Got CatchUps", "catch_ups", catch_ups)

	// Mon Jan 2 15:04:05 -0700 MST 2006
	d1 := "2018-03-10"
	startDate, err := time.Parse("2006-01-02", d1)
	if err != nil {
		panic(fmt.Sprint("failed to parse ", d1))
	}
	log.Info("Parsed startDate", "startDate", startDate)

	d2 := "2018-03-17"
	endDate, err := time.Parse("2006-01-02", d2)
	if err != nil {
		panic(fmt.Sprint("failed to parse ", d2))
	}
	log.Info("Parsed endDate", "endDate", endDate)
	/*
		options := getOptions(startDate, endDate)

		c1 := CatchUp{
			Name:      "First catchup",
			Details:   "First attempt at a catchup record",
			StartDate: startDate,
			EndDate:   endDate,
			Options:   options,
		}

		db.Create(&c1)
	*/
}

/*
func getOptions(startDate time.Time, endDate time.Time) (options []Option) {
	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}
	for day := startDate; day.Before(endDate) || day.Equal(endDate); day = day.AddDate(0, 0, 1) {
		o := Option{Date: day}
		options = append(options, o)
	}
	return options
}
*/
