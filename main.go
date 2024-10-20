package main

import "github.com/gin-gonic/gin"

type ID struct {
	Username string
	Password string
}

func main() {
	r := gin.Default()
	r.POST("/signup", IDHandler1)
	r.Run()
}

func IDHandler1(c *gin.Context) {
	var request ID
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	// user1 := ID{"Anirudh", "ANI123"}
	// user2 := ID{"Abhinav", "ABHI123"}

}
