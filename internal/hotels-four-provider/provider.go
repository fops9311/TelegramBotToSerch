package hotelsfourprovider

import "TelegramBotToSerch/internal/model"

type HotelFourProvider struct {
}

func NewHotelFourProvider() *HotelFourProvider {
	return &HotelFourProvider{}
}

func (p HotelFourProvider) GetLocationHotels(lat float64, lon float64) []model.Hotel {
	return []model.Hotel{
		{
			Name: "Bullshit",
		},
	}
}
