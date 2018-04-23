package main

//go:generate sqlboiler mysql --wipe

import (
	"database/sql"
	"fmt"
	"github.com/d5ve/catch-up/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	router.GET("/new", func(c *gin.Context, db) {})
	router.POST("/new", func(c *gin.Context, db) {})
}

func createCatchup(db *sql.DB) {
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

	optionDates := getDates(startDate, endDate)
	log.Info("Generated option dates", "dates", optionDates)

	catch_up := models.CatchUp{
		Name:      "My first catchup",
		StartDate: startDate,
		EndDate:   endDate,
	}

	err = catch_up.Insert(db)
	if err != nil {
		panic(fmt.Sprint("Failed to INSERT catch_up", err))
	}

	for _, oDate := range optionDates {
		option := models.Option{
			CatchUpID: catch_up.ID,
			Date:      oDate,
		}
		err := option.Insert(db)
		if err != nil {
			panic(fmt.Sprint("Failed to INSERT option", err))
		}
	}

	catch_ups, err := models.CatchUps(db).All()
	log.Info("Got CatchUps", "catch_ups", spew.Sdump(catch_ups))

	options, err := models.Options(db).All()
	log.Info("Got CatchUps", "options", spew.Sdump(options))

}

func getDates(startDate time.Time, endDate time.Time) (optionDates []time.Time) {
	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}
	for day := startDate; day.Before(endDate) || day.Equal(endDate); day = day.AddDate(0, 0, 1) {
		optionDates = append(optionDates, day)
	}
	return optionDates
}
