package api

import (
	"github.com/gin-gonic/gin"
	"FloorPlanXpert/internal/handlers"
	"net/http"
)

func SetupRouter() *gin.Engine {
	// Define routes
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

	r.GET("/", handlers.HomePage)
    r.GET("/user/login-register", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login_register.html", nil)
    })
	r.POST("/user/register", handlers.RegisterUser)
	r.POST("/user/login", handlers.LoginUser)
	r.POST("/rooms/create", handlers.CreateRoom)
	r.POST("/rooms/book", handlers.BookRoom)
	return r
}