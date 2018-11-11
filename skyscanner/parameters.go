package skyscanner

// Parameters generic request type for skyscanner api
type Parameters struct {
	Country          string `json:"country"`
	Currency         string `json:"currency"`
	Locale           string `json:"locale"`
	OriginPlace      string `json:"originPlace"`
	DestinationPlace string `json:"destinationPlace"`
	OutbandDate      string `json:"outboundDate"`
	Adults           int    `json:"adults"`
	InboundDate      string `json:"inboundDate"`
}
