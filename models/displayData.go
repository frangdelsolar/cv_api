package models

type DisplayData struct {
	Model      `gorm:"embedded"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	City       string `json:"city"`
	Country    string `json:"country"`
	ProfilePic string `json:"profile_pic"`
}

type DisplayDatas []DisplayData
