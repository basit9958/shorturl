package models

import "time"

type Request struct {
	URL        string `json:"url"`
	ShortUrl   string `json:"short"`
	Expiration time.Duration
}

type Response struct {
	URL        string `json:"url"`
	ShortUrl   string `json:"short"`
	Expiration time.Duration
}
