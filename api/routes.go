package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"FloorPlanXpert/internal/handlers"
	"FloorPlanXpert/internal/utils"
	"FloorPlanXpert/internal/models"
)

func SetupRouter() *gin.Engine {
	// Define routes
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

	r.GET("/", handlers.HomePage)
    r.GET("/user/login-register", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login_register.html", nil)
    })
	r.POST("/user/register", func(c *gin.Context) {
        // Handle register in a Goroutine
        user := models.User{} // Initialize user struct
		if err := c.ShouldBind(&user); err != nil {
			utils.LogError(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resultChan := make(chan bool)

		// Call RegisterUser in a Goroutine
		go handlers.RegisterUser(c, user, resultChan)

		// Wait for the result from RegisterUser

		result := <-resultChan

		// Handle the result accordingly
		if !result {
			// Handle the scenario where the registration failed
			utils.Log("Registration failed")
			return
		}
    })
    r.POST("/user/login", func(c *gin.Context) {
		// Pass a copy of the context and any required data to the Goroutine
		user := models.User{} // Initialize user struct
		if err := c.ShouldBind(&user); err != nil {
			utils.LogError(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resultChan := make(chan bool)

		// Call LoginUser in a Goroutine
		go handlers.LoginUser(c, user, resultChan)

		// Wait for the result from LoginUser
		result := <-resultChan

		// Handle the result accordingly
		if !result {
			// Handle the scenario where the login failed
			utils.Log("Login failed")
			return
		}
	})
	
	r.POST("/rooms/create", handlers.CreateRoom)
	r.POST("/rooms/book", handlers.BookRoom)

	return r
}