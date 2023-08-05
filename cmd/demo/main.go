package main

import (
	"gin-mall/internal/queue"
	"github.com/hibiken/asynq"
	"log"
)

func main() {
	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		})

	mux := asynq.NewServeMux()
	mux.HandleFunc(queue.TypeEmailDelivery, queue.HandleEmailDeliveryTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("cant run server %v", err)
	}
}
