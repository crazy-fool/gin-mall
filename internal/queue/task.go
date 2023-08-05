package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

const TypeEmailDelivery = "email:deliver"
const TypeImageResize = "image:resize"

type EmailDeliveryPayload struct {
	UserId     int
	TemplateID string
}

type ImageResizePayload struct {
	SourceUrl string
}

func NewEmailDeliveryTask(userID int, tplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserId: userID, TemplateID: tplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func NewImageResizeTask(src string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageResizePayload{SourceUrl: src})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeImageResize, payload), nil

}

func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("errrrrrrr,%v:%w", err, asynq.SkipRetry)
	}

	log.Printf("userID %v,teplate_id %v", p.UserId, p.TemplateID)
	return nil
}

type ImageProcessor struct {
}

func (i *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("errrrrrrr,%v:%w", err, asynq.SkipRetry)
	}

	log.Printf("source %v", p.SourceUrl)
	return nil
}

func NewImageprocessor() *ImageProcessor {
	return &ImageProcessor{}
}
