package models

type Request struct {
	URL      string `json:"url"`
	ShortUrl string `json:"short"`
}

type Response struct {
	URL      string `json:"url"`
	ShortUrl string `json:"short"`
}
