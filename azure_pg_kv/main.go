package main

import (
	"fmt"
	"net/http"
	"os"

	"azure_pg_kv/dbutils"

	"github.com/gin-gonic/gin"
)

var db *dbutils.Database

func resolvePort() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func Ping(c *gin.Context) {
	if err := db.Ping(); err != nil {
		message := fmt.Sprintf("Database ping failed: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}

func main() {

	db = dbutils.NewConnection()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.GET("/api/ping", Ping)
	r.Run(resolvePort())
}
