package domain

type AdviceBuySell struct {
    NiftyIdentifier string  `json:"niftyIdentifier"`
    Index           int     `json:"index"`
    Advice          string  `json:"advice"`
    TargetPrice     float64 `json:"targetPrice"`
    SMA200Support   float64 `json:"sma200Support"`
}

type CandleData struct {
    Date   string  `json:"date"`
    Open   float64 `json:"open"`
    High   float64 `json:"high"`
    Low    float64 `json:"low"`
    Close  float64 `json:"close"`
    Volume float64 `json:"volume"`
}
