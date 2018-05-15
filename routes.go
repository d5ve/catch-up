package main

import (
	"database/sql"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func (env *Env) getIndex(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("count", 1)
	session.Save()
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

	routes := gin.Default()

	store := cookie.NewStore([]byte("3@ZN2@PDxLBEq#AH7wMAf%ij$U59U%Tg"))
	routes.Use(sessions.Sessions("catch-up-session", store))

	routes.GET("/", env.getIndex)
	routes.GET("/new", env.getNew)
	routes.POST("/new", env.postNew)
	routes.GET("/edit", env.getEdit)
	routes.POST("/edit", env.postEdit)
	routes.GET("/vote", env.getVote)
	routes.POST("/vote", env.postVote)
	routes.Static("/static", "./static")

	routes.LoadHTMLGlob("templates/*.html")

	return routes
}
