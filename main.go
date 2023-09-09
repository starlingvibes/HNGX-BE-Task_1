package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		slackName := c.DefaultQuery("slack_name", "Chidera_Anichebe")
		track := c.DefaultQuery("track", "Backend")

		c.JSON(200, gin.H{
			"slack_name": slackName,
			"current_day": time.Now().Format("Monday"),
			"utc_time": time.Now().UTC().Format(time.RFC3339),
			"track": track,
			"github_file_url": "https://github.com/starlingvibes/HNGX-BE-Task_1/blob/main/main.go",
			"github_repo_url": "https://github.com/starlingvibes/HNGX-BE-Task_1",
			"status_code": 200,
		})
	})
	 // listen and serve on 0.0.0.0:3000
	r.Run(":3000")
}