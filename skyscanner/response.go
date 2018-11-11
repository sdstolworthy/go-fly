package skyscanner

import "errors"

// Response generic response from skyscanner api
type Response struct {
	Quotes []struct {
		QuoteID     int     `json:"QuoteId"`
		MinPrice    float64 `json:"MinPrice"`
		Direct      bool    `json:"Direct"`
		OutboundLeg struct {
			CarrierIds    []int  `json:"CarrierIds"`
			OriginID      int    `json:"OriginId"`
			DestinationID int    `json:"DestinationId"`
			DepartureDate string `json:"DepartureDate"`
		} `json:"OutboundLeg"`
		InboundLeg struct {
			CarrierIds    []int  `json:"CarrierIds"`
			OriginID      int    `json:"OriginId"`
			DestinationID int    `json:"DestinationId"`
			DepartureDate string `json:"DepartureDate"`
		} `json:"InboundLeg"`
		QuoteDateTime string `json:"QuoteDateTime"`
	} `json:"Quotes"`
	Places []struct {
		PlaceID        int    `json:"PlaceId"`
		IataCode       string `json:"IataCode"`
		Name           string `json:"Name"`
		Type           string `json:"Type"`
		SkyscannerCode string `json:"SkyscannerCode"`
		CityName       string `json:"CityName"`
		CityID         string `json:"CityId"`
		CountryName    string `json:"CountryName"`
	} `json:"Places"`
	Carriers []struct {
		CarrierID int    `json:"CarrierId"`
		Name      string `json:"Name"`
	} `json:"Carriers"`
	Currencies []struct {
		Code                        string `json:"Code"`
		Symbol                      string `json:"Symbol"`
		ThousandsSeparator          string `json:"ThousandsSeparator"`
		DecimalSeparator            string `json:"DecimalSeparator"`
		SymbolOnLeft                bool   `json:"SymbolOnLeft"`
		SpaceBetweenAmountAndSymbol bool   `json:"SpaceBetweenAmountAndSymbol"`
		RoundingCoefficient         int    `json:"RoundingCoefficient"`
		DecimalDigits               int    `json:"DecimalDigits"`
	} `json:"Currencies"`
}

// Prices returns an array of the lowest prices for a route and date
func (r *Response) Prices() []float64 {
	priceList := make([]float64, len(r.Quotes))
	for i, v := range r.Quotes {
		priceList[i] = v.MinPrice
	}
	return priceList
}

// A QuoteSummary is a summary of a outbound trip with its price and date
type QuoteSummary struct {
	Price         float64
	DepartureDate string
	InboundDate   string
}

// LowestPrice prints a list of the lowest prices and their accompanying dates
func (r *Response) LowestPrice() (*QuoteSummary, error) {
	quote := QuoteSummary{
		99999999,
		"",
		"",
	}
	for _, v := range r.Quotes {
		if v.MinPrice < quote.Price {
			quote.Price = v.MinPrice
			quote.DepartureDate = v.OutboundLeg.DepartureDate
			quote.InboundDate = v.InboundLeg.DepartureDate
		}
	}
	if quote.Price == 99999999 {
		return nil, errors.New("No quote found")
	}
	return &quote, nil

}
