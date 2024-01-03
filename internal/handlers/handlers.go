package handlers

import (
	"FloorPlanXpert/internal/db"
	"FloorPlanXpert/internal/models"
	"FloorPlanXpert/internal/utils"
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
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

    existingUser, err := db.GetUserByUsername(newUser.Username)
    if err != nil {
        utils.LogError(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!", "userID": existingUser.UserID})
}

func LoginUser(c *gin.Context) {
    utils.Log("Received a request to login a user.")
    // Decode the request body to extract user login details
    var user models.User
    if err := c.ShouldBind(&user); err != nil {
        utils.LogError(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    fmt.Printf("Received User Data: %+v\n", user)
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

    c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully!", "userID": existingUser.UserID})
}


func CreateRoom(c *gin.Context) {
    utils.Log("Received a request to create a room.")
    // Decode the request body to extract room details
    var newRoom models.Room

    newRoom.BookingUserID = 0;
    newRoom.Lastbookingendtime = time.Time{}

    if err := c.ShouldBindJSON(&newRoom); err != nil {
        utils.LogError(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    fmt.Println(newRoom)
    if newRoom.RoomName != 0 && newRoom.Capacity != 0 {
        if err := db.InsertRoom(&newRoom); err != nil {
            utils.LogError(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room details"})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"message": "Room created successfully!"})
}

func BookRoom(c *gin.Context) {
    var req models.BookRoomRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    currentTime := time.Now()

    // Fetch rooms matching the criteria from the database
    var rooms []models.Room
    err := db.DB.Model(&rooms).
        Where("capacity >= ?", req.RequiredCapacity).
        Where("lastbookingendtime IS NULL OR lastbookingendtime < ?", currentTime).
        Order("capacity ASC").
        Select()

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rooms"})
        return
    }

    // Check if there are available rooms
    if len(rooms) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No available rooms matching the criteria"})
        return
    }

    // Update the first room from the fetched list
    roomToUpdate := rooms[0]
    roomToUpdate.Lastbookingendtime = currentTime.Add(time.Minute * time.Duration(req.Duration))
    roomToUpdate.BookingUserID = req.UserID

    _, err = db.DB.Model(&roomToUpdate).
        Where("room_id = ?", roomToUpdate.RoomID).
        Column("lastbookingendtime", "booking_user_id").
        Update()

    if err != nil {
        utils.LogError(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book the room"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Room booked successfully", "roomName": roomToUpdate.RoomName})
}
