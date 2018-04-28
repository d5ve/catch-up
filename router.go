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

func (env *Env) getNew(c *gin.Context) {
	log.Info("Got DB handle", "db", env.db)
	c.HTML(
		http.StatusOK,
		"new.html",
		gin.H{
			"title": "New catchup",
		},
	)
}

func (env *Env) postNew(c *gin.Context) {}

func (env *Env) getEdit(c *gin.Context)  {}
func (env *Env) postEdit(c *gin.Context) {}

func (env *Env) getVote(c *gin.Context)  {}
func (env *Env) postVote(c *gin.Context) {}

func setupRouter(db *sql.DB) *gin.Engine {

	env := &Env{db: db}

	router := gin.Default()

	router.GET("/new", env.getNew)
	router.POST("/new", env.postNew)
	router.GET("/edit", env.getEdit)
	router.POST("/edit", env.postEdit)
	router.GET("/vote", env.getVote)
	router.POST("/vote", env.postVote)

	router.LoadHTMLGlob("templates/*")

	return router
}
