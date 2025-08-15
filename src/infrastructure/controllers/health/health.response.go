package controllers

type HealthResponse struct {
	Author  string `json:"author"`
	Message string `json:"message"`
	Counter int    `json:"counter"`
}
