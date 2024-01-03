package utils

import (
    "log"
    "os"
)

func Log(message string) {
    log.Println(message)
}

func LogError(err error) {
    log.Println("Error:", err.Error())
}

func InitLogging() {
    // Example: Initialize log settings if needed (log file, log format, etc.)
    log.SetOutput(os.Stdout)
    log.SetFlags(log.LstdFlags | log.Lshortfile)
}