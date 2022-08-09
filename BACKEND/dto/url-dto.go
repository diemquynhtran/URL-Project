package dto


type CreateUrl struct {
	LongURL string `json:"long" form:"long" binding:"required"`
	UserID  uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type CreateFreeUrl struct {
	LongURL string `json:"long" form:"long" binding:"required"`
}

type UrlResponse struct {
	LongURL     string `json:"long"`
	ShortURL    string `json:"short" gorm:"UNIQUE"`
	SnowflakeId uint64 `json:"snowflake" gorm:"UNIQUE"`
	NumberClick uint16 `json:"count"`
	Status 		bool `json:"status" gorm:"default:true"`
}

type GetUrlByName struct {
	ShortURL string `json:"short" gorm:"UNIQUE"`
}

type GetUrlByUSer struct {
	ShortURL string `json:"short" gorm:"UNIQUE"`
	UserId   uint64 `json:"user_id"`
}
