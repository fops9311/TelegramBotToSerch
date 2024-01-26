package telegrambotsearchservice

import (
	"TelegramBotToSerch/internal/model"
	"sync"

	"go.uber.org/zap"
)

type HotelInfoProvider interface {
	GetLocationHotels(lat float64, lon float64) []model.Hotel
}

type Requester interface {
	Requests() chan model.HotelsRequest
}

type Service struct {
	logger        *zap.Logger
	hotelProvoder HotelInfoProvider
	requesters    []Requester
}

type Opts struct {
	Logger   *zap.Logger
	Provider HotelInfoProvider
}

func NewService(opts Opts) *Service {
	return &Service{
		logger:        opts.Logger,
		hotelProvoder: opts.Provider,
	}
}

func (s *Service) Run() {
	var wg sync.WaitGroup
	var req chan model.HotelsRequest = make(chan model.HotelsRequest)
	for _, requester := range s.requesters {
		wg.Add(1)
		go func(requester Requester) {
			for r := range requester.Requests() {
				s.logger.Info("got new request")
				req <- r
			}
			wg.Done()
		}(requester)
	}
	var done chan struct{} = make(chan struct{})
	go func(done chan struct{}) {
		for r := range req {
			s.hotelProvoder.GetLocationHotels(r.Lat, r.Lon)
			s.logger.Info("got provider response")
		}
		close(done)
	}(done)
	wg.Wait()
	close(req)
	<-done
}

func (s *Service) WithRequester(r Requester) *Service {
	s.requesters = append(s.requesters, r)
	return s
}
