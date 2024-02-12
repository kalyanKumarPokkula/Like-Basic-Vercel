package controllers

import (
	"fmt"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/kalyanKumarPokkula/vercel/helpers"
)

type repourl struct{
	Url string `json:"url"`
}


func Deploy(c *gin.Context) {
	str, err := helpers.GenerateRandomString()
	var body repourl
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to read body",
			"success": "false",
		})

		return
	}
	fmt.Println(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to generate random string",
		})
		return
	}
	dir := "../output/" + str
	gitrepo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      body.Url,
		Progress: os.Stdout,
	})

	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to download github repo",
		})

		return
	}

	folderPath := "../output/" + str

	upload_err := helpers.UploadFolder(folderPath, str)
	
	if upload_err != nil {
		c.JSON(http.StatusBadRequest ,gin.H{
			"message" : "Failed to upload files to s3 bucket",
		})
	}

	redisQueue_err := helpers.RedisQueue(str)

	if redisQueue_err != nil {
		c.JSON(http.StatusBadRequest ,gin.H{
			"message" : "Failed to push to redis queue",
		})
	}

	fmt.Println(gitrepo)
	c.JSON(200, gin.H{
		"message": "Successfully downloaded repo",
		"userID" : str,
	})
}