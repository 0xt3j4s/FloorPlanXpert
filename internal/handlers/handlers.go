package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "FloorPlanXpert/internal/db"
    "FloorPlanXpert/internal/models"
    "FloorPlanXpert/internal/utils"
)


func HomePage(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
}


func RegisterUser(c *gin.Context) {
    utils.Log("Received a request to register a user.")
    // Decode the request body to extract user registration details
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        utils.LogError(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := db.GetUserByUsername(newUser.Username)
    if err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
        return
    }

    // Insert new user into the database
    if err := db.InsertUser(&newUser); err != nil {
        utils.LogError(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

func LoginUser(c *gin.Context) {
    utils.Log("Received a request to login a user.")
    // Decode the request body to extract user login details
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.LogError(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Retrieve user from the database
    existingUser, err := db.GetUserByUsername(user.Username)
    if err != nil {
        utils.LogError(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
        return
    }

    // Check if the password is correct
    if user.Password != existingUser.Password {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully!"})
}