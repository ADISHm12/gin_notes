package main

import (
	"gin_notes/controllers"
	"gin_notes/middlewares"
	"gin_notes/models"
	"log"
	"net/http"
	controllers_helpers "gin_notes/controllers/helpers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	
	r.LoadHTMLGlob("templates/**/**")
	r.Static("/vendor", "./static/vendor")
	models.ConnectDatabase()
	models.DBMigrate()
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes",store))
	r.Use(middlewares.AuthenticateUser())
	
	
	notes := r.Group("/notes")
	{
		notes.GET("/",controllers.NotesIndex)
		notes.GET("/new",controllers.NotesNew)
		notes.POST("/",controllers.NotesCreate)
		notes.GET("/:id",controllers.NotesShow)
		notes.GET("/edit/:id",controllers.NotesEditPage)
		notes.POST("/:id",controllers.NotesUpdate)
		notes.DELETE("/:id",controllers.NotesDelete)
	
	}
	
	r.GET("/login",controllers.LoginPage)
	r.GET("signup",controllers.SignupPage)
	r.POST("/login",controllers.Login)
	r.POST("/signup",controllers.Signup)
	r.POST("/logout",controllers.Logout)


	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,"home/index.html",controllers_helpers.SetPayload(c,gin.H{
			
			"title":"notes application",
			"logged_in" : controllers_helpers.IsUserLoggedIn(c),
		}))
	})
	
	log.Println("server started")
	
	r.Run()
}