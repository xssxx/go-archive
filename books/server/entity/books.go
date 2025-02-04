package entity

type Book struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	Title         string  `json:"title" gorm:"type:varchar(255);not null"`
	Author        string  `json:"author" gorm:"type:varchar(255);not null"`
	PublishedDate *string `json:"published_date" gorm:"type:date"`
	Genre         *string `json:"genre" gorm:"type:varchar(100)"`
	Price         float64 `json:"price" gorm:"type:numeric(10,2);not null"`
	Stock         int     `json:"stock" gorm:"default:0"`
	CreatedAt     *string `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt     *string `json:"updated_at" gorm:"type:timestamp"`
}
