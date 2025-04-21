package models

type Match struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Home  string `json:"home"`
	Away  string `json:"away"`
	Score string `json:"score"`
}
