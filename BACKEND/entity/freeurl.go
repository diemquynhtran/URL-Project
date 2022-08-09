package entity

type FreeUrl struct {
	LongURL     string `json:"long"`
	ShortURL    string `json:"short" gorm:"UNIQUE"`
	SnowflakeId uint64 `json:"snowflake" gorm:"primary_key; UNIQUE"`
}

