package dto


type CreateUrl struct {
	LongURL string `json:"long" form:"long" binding:"required"`
	UserID  uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type UrlResponse struct {
	LongURL     string `json:"long"`
	ShortURL    string `json:"short" gorm:"UNIQUE"`
	SnowflakeID uint64 `json:"snowflake" gorm:"primaryKey; UNIQUE"`
	NumberClick uint16 `json:"count"`
	UserId      uint64 `json:"user_id"`
}

type GetUrlByName struct {
	ShortURL string `json:"short" gorm:"UNIQUE"`
}

type GetUrlByUSer struct {
	ShortURL string `json:"short" gorm:"UNIQUE"`
	UserId   uint64 `json:"user_id"`
}
