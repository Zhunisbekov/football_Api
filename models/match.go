package models

type Match struct {
	ID    int    `json:"id"`
	Home  string `json:"home"`
	Away  string `json:"away"`
	Score string `json:"score"`
}
