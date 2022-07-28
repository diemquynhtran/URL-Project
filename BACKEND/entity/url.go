package entity

type Url struct {
	LongURL     string `json:"long"`
	ShortURL    string `json:"short" gorm:"UNIQUE"`
	SnowflakeId uint64 `json:"snowflake" gorm:"primaryKey; UNIQUE"`
	NumberClick uint16 `json:"click"`
	Uid         uint64 `json:"user_id"`
}
