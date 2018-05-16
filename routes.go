package main

import (
	"database/sql"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
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

type CatchUp struct {
	Name     string    `form:"name" binding:"required"`
	Details  string    `form:"details"`
	StartDay time.Time `form:"start_day" binding:"required,futuredate" time_format:"02/01/2006"`
	EndDay   time.Time `form:"end_day" binding:"required,gtfield=StartDay" time_format:"02/01/2006"`
}

func futureDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func (env *Env) postNew(c *gin.Context) {

	var cu CatchUp
	if err := c.ShouldBind(&cu); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("futuredate", futureDate)
	}

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
