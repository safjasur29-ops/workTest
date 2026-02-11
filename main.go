package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("file create error", err.Error())
		return
	}

	logger := &lumberjack.Logger{
		Filename:   f.Name(),
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     28,   // days
		Compress:   true, // disabled by default
	}

	gin.DefaultWriter = io.MultiWriter(logger, os.Stdout)
	log.SetOutput(logger)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
		log.Println("Work good")
	})
	log.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
