package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func (env *Env) getIndex(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"title": "Catch-up"},
	)
}

func (env *Env) getNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"new.html",
		gin.H{
			"title": "New catchup",
		},
	)
}

func (env *Env) postNew(c *gin.Context) {
	log.Info("Got DB handle", "db", env.db)

}

func (env *Env) getEdit(c *gin.Context)  {}
func (env *Env) postEdit(c *gin.Context) {}

func (env *Env) getVote(c *gin.Context)  {}
func (env *Env) postVote(c *gin.Context) {}

func setupRoutes(db *sql.DB) *gin.Engine {

	env := &Env{db: db}

	router := gin.Default()

	router.GET("/", env.getIndex)
	router.GET("/new", env.getNew)
	router.POST("/new", env.postNew)
	router.GET("/edit", env.getEdit)
	router.POST("/edit", env.postEdit)
	router.GET("/vote", env.getVote)
	router.POST("/vote", env.postVote)
	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*.html")

	return router
}