package basic

import (
	"encoding/json"
	"fmt"
	"github.com/herhe-com/framework/facades"
	constant "project.io/shop/constants/queue"
)

func PublishError(data []byte, queue string, err error) {

	message, _ := json.Marshal(constant.BasicErrorMessage{
		Queue:   queue,
		Message: string(data),
		Error:   fmt.Sprintf("%v", err),
	})

	_ = facades.Queue.Producer(message, constant.BasicError, []string{constant.BasicError})
}

func Publish(data []byte, queue string) error {
	return facades.Queue.Producer(data, queue, []string{queue})
}
