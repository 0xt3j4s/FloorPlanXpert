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

func RegisterUser(c *gin.Context, newUser models.User, resultChan chan<- bool) {
    // Make two channels to receive the user data and errors
    existingUserCh := make(chan *models.User)
    errorCh := make(chan error)
    errorCh2 := make(chan error)

    go func()  {
        existingUser, err := db.GetUserByUsername(newUser.Username)
        if err == nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
            existingUserCh <- existingUser
            return
        } else {
            errorCh <- err
            return
        }
    }()


    // Wait for either the existingUser or an error
    select {
    case existingUser := <-existingUserCh:
        c.JSON(http.StatusBadRequest, gin.H{"error": existingUser.Username + " already exists"})
        resultChan <- false // Indicate failure through the channel
        return

    case err := <-errorCh: 
        utils.Log(err)
        go func ()  {
            // wait for insert user
            if err := db.InsertUser(&newUser); err != nil {
                errorCh2 <- err
                return
            }
        }()
        // if no error return successful
        select {
        case err := <-errorCh2:
            utils.LogError(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            resultChan <- false // Indicate failure through the channel
            return
        default:
            c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!", "userID": newUser.UserID})
            resultChan <- true // Indicate success through the channel
            return
        }
    }
}


func LoginUser(c *gin.Context, user models.User, resultChan chan<- bool) {
    existingUserCh := make(chan *models.User) // Channel to receive the user data
    errorCh := make(chan error) // Channel to receive errors

    // Perform the database query within a Goroutine
    go func() {
        existingUser, err := db.GetUserByUsername(user.Username)
        if err != nil {
            errorCh <- err
            return
        }
        existingUserCh <- existingUser
    }()

    // Wait for either the existingUser or an error
    select {
    case existingUser := <-existingUserCh:
        // Check password and proceed accordingly
        if user.Password != existingUser.Password {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
            resultChan <- false // Indicate failure through the channel
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully!", "userID": existingUser.UserID})
        resultChan <- true // Indicate success through the channel

    case err := <-errorCh:
        utils.LogError(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
        resultChan <- false // Indicate failure through the channel
        return
    }
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