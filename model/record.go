package model

type Record struct {
	Id        int    `gorm:"type:int" json:"record_id"`
	ChainAddr string `gorm:"size:255,not null,unique" json:"chain_addr"`
	Reward    int    `gorm:"type:int，not null" json:"reward_id"`
}

func (Record) TableName() string {
	return "record"
}
