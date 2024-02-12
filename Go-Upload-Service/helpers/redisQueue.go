package helpers

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)


func RedisQueue(task string) error {
    // Connect to Redis
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    // Push task into the queue
    err := rdb.LPush(context.Background(), "task_queue", task).Err()
    if err != nil {
        // panic(err)
		return err
    }
    fmt.Println("Task pushed into the queue")

	return nil

    // Continuously monitor the queue for tasks
    // for {
    //     // Pop task from the queue (blocking operation)
    //     taskJSON, err := rdb.BRPop(context.Background(), 0, "task_queue").Result()
    //     if err != nil {
    //         panic(err)
    //     }

    //     // Process the task
    //     var task Task
    //     if err := json.Unmarshal([]byte(taskJSON[1]), &task); err != nil {
    //         panic(err)
    //     }
    //     fmt.Println("Processing task:", task)
    //     // Perform your task processing logic here
    // }
}
