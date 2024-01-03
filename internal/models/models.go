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
    lastBookingEndTime time.Time
    Capacity      int
}

type Booking struct {
    UserID         int
    RoomID         int
    StartTime  time.Time
    EndTime    time.Time
}