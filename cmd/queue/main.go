package main

import (
	"gin-mall/internal/queue"
	"github.com/hibiken/asynq"
	"log"
	"sync"
)

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "146.56.205.28:6379"})
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			log.Println("error", err)
		}
	}(client)
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(id int) {
			t1, err := queue.NewEmailDeliveryTask(id, "test tpl")
			if err != nil {
				log.Fatalf("创建任务失败 %v", err)
			}

			info, err := client.Enqueue(t1)
			if err != nil {
				log.Fatalf("发送任务失败 %v", err)
			}

			log.Printf("id:%v;taskid:%s;queue:%s", id, info.ID, info.Queue)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
