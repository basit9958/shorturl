package models

type request struct {
	URL   string `json:"url"`
	Short string `json:"short"`
}

type response struct {
	URL   string `json:"url"`
	SHORT string `json:"short"`
}
