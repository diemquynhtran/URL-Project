package entity

type Url struct {
	LongURL     string `json:"long"`
	ShortURL    string `json:"short" gorm:"UNIQUE"`
	SnowflakeId uint64 `json:"snowflake" gorm:"primary_key; UNIQUE"`
	NumberClick uint16 `json:"click"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE" json:"user"`
	Status 		bool `json:"status" gorm:"default:true"`
}

