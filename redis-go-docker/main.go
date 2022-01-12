package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// go get -u github.com/go-redis/redis
// If you encounter problems like I did with a newer version of Go. Run the following:
// GO111MODULE="off" go get -u github.com/go-redis/redis and restart the IDE.
// Restart IDE

func main() {
	fmt.Println("Hello everyone!")

	// Setup the connection to redis. "redis" is the name of the container which lets
	// docker handle the networking. "mypassword" is the password used in docker-compose.yml.
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "mypassword",
		DB:       0,
	})

	// Set a key and value for testing.
	time.Sleep(5 * time.Second)
	err := client.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("key is set")
	// For testing purposes sleep for 10 mins to keep container alive. Should serve as a web app.
	time.Sleep(10 * time.Minute)

}
