package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/kalyanKumarPokkula/vercel-deploy/helpers"
)

func main() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Continuously monitor the queue for tasks
    for {
        // Pop task from the queue (blocking operation)
        task, err := rdb.BRPop(context.Background(), 0, "task_queue").Result()
        if err != nil {
            panic(err)
        }
        // download the folder from s3 bucket
		helpers.S3Download(task[1])

		// build the project
		dir := "../output/" + task[1]
		helpers.BuildProject(dir)

		// upload build folder to s3 bucket
		folderPath := "../output/"+task[1]+"/dist"
		helpers.UploadBuildFolder(folderPath , task[1])
        fmt.Println("Processing task:", task)
        // Perform your task processing logic here
    }

}