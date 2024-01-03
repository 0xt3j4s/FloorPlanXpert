package db

import (
    "fmt"
    "github.com/go-pg/pg/v10"
    "FloorPlanXpert/internal/models"
    "FloorPlanXpert/internal/utils"
)

var DB *pg.DB

func Connect() error {
    DB = pg.Connect(&pg.Options{
        User:     "moveinsync",
        Password: "moveinsync",
        Database: "floorplan",
        Addr:     "localhost:5432",
    })

    _, err := DB.Exec("SELECT 1")
    if err != nil {
        return fmt.Errorf("failed to connect to the database: %w", err)
    }

    fmt.Println("Connected to the database")

    
    return nil
}


func Close() {
    DB.Close()
}


// Insert a new user into the database
func InsertUser(newUser *models.User) error {
    _, err := DB.Model(newUser).Returning("*").Insert()
    return err
}

func InsertRoom(newRoom *models.Room) error {
    utils.Log("Inserting new room")
    _, err := DB.Model(newRoom).Returning("*").Insert()
    return err
}

func InsertBooking(newBooking *models.Booking) error {
    utils.Log("Inserting new booking")
    _, err := DB.Model(newBooking).Returning("*").Insert()
    return err
}

// Retrieve a user by username
func GetUserByUsername(username string) (*models.User, error) {
    user := &models.User{}
    err := DB.Model(user).Where("username = ?", username).Select()
    if err != nil {
        return nil, err
    }
    return user, nil
}