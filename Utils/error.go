package utils

type DefaultErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
