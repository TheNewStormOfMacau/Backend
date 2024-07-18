package model

type User struct {
	Id        int    `gorm:"type:int" json:"user_id"`
	ChainAddr string `gorm:"size:255,not null,unique" json:"chain_addr"`
	Total     int    `gorm:"type:int" json:"total"`
	Balance   int    `gorm:"type:int" json:"balance"`
}

func (User) TableName() string {
	return "user"
}
