package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"api-golang/database"
)

func init() {
	errDB := database.InitDB(os.Getenv("DATABASE_URL"))
	if errDB != nil {
		log.Fatalf("⛔ Unable to connect to database: %v\n", errDB)
	} else {
		log.Println("DATABASE CONNECTED 🥇")
	}

}

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		// AllowOrigins:  []string{"http://127.0.0.1:5173/"},
		AllowMethods:  []string{"GET"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
			// return origin == "http://127.0.0.1:8080/"
		},
		// MaxAge: 12 * time.Hour,
	}))
	var tm time.Time

	r.GET("/", func(c *gin.Context) {
		tm = database.GetTime(c)
		c.JSON(200, gin.H{
			"api": "golang",
			"now": tm,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
