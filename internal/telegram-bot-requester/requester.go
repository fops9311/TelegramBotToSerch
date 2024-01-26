package telegrambotrequester

import (
	"TelegramBotToSerch/internal/model"
	"time"
)

type TelegramBotRequester struct {
}

func (r *TelegramBotRequester) Requests() chan model.HotelsRequest {
	var requests chan model.HotelsRequest = make(chan model.HotelsRequest)
	go func() {
		defer close(requests)
		for {
			requests <- model.HotelsRequest{Lat: 10, Lon: 15}
			time.Sleep(time.Second)
		}
	}()
	return requests
}
