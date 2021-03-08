package models

type AddRequest struct {
	Number float64 `json:"number"`
}

type AddResponse struct {
	Message string `json:"message"`
}

type CalcResponse struct {
	Sum     float64 `json:"sum"`
	Average float64 `json:"average"`
	Count   int     `json:"count"`
}
