package dto

type SendRequest struct {
	To      string  `json:"id"`
	Balance float64 `json:"balance"`
}
