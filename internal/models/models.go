package models

import (
    "time"
)


type User struct {
    UserID   int `pg:"user_id,pk"`
    Level    int
    Name    string
    Username string
    Password string
}

type Room struct {
    RoomID        int
    RoomName      int
    BookingUserID  int
    Capacity      int
    Lastbookingendtime time.Time `pg:"lastbookingendtime"`
}

type BookRoomRequest struct {
    UserID          int `json:"userID"`
    RequiredCapacity int `json:"requiredCapacity"`
    Duration int `json:"duration"`
}

type Booking struct {
    UserID         int
    RoomID         int
    StartTime  time.Time
    EndTime    time.Time
}