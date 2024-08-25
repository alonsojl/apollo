package models

import (
	"net/http"
	"time"
)

type TemplateResponse struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Datetime  string `json:"datetime"`
	Timestamp int64  `json:"timestamp"`
}

func NewTemplateResponse() TemplateResponse {
	return TemplateResponse{
		Status:    "success",
		Code:      http.StatusOK,
		Datetime:  time.Now().Format(time.DateTime),
		Timestamp: time.Now().Unix(),
	}
}
