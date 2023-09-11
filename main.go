package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type student struct {
	SlackName     string `json:"slack_name"`
	CurrentDay    string `json:"current_day"`
	UTCTime       string `json:"utc_time"`
	Track         string `json:"track"`
	GithubFileURL string `json:"github_file_url"`
	GithubRepoURL string `json:"github_repo_url"`
	StatusCode    int    `json:"status_code"`
}

func main() {
	port := "3000"
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server works")
	})
	r.GET("/api", doSomething)
	r.Run(":" + port)
}

func doSomething(c *gin.Context) {
	sName := strings.ToLower(c.Query("slack_name"))
	track := strings.ToLower(c.Query("track"))

	fmt.Println("sName", sName)
	fmt.Println("track", track)

	var st student
	{
		st.SlackName = "meshach"
		st.Track = "backend"
		st.GithubFileURL = "https://github.com/meshachdamilare/hng_simple_server/blob/main/main.go"
		st.GithubRepoURL = "https://github.com/meshachdamilare/hng_simple_server"

	}

	if sName == st.SlackName && track == st.Track {
		st.CurrentDay = time.Now().Weekday().String()
		st.UTCTime = time.Now().UTC().Format(time.RFC3339)
		st.StatusCode = http.StatusOK
		c.JSON(http.StatusOK, st)
		return
	}
	st.StatusCode = http.StatusBadRequest
	c.JSON(http.StatusBadRequest, gin.H{"error": "request not found", "st": st.StatusCode})
}
