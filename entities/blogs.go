package entities

import (
	"time"
)

type Blogs struct {
	Id        int       `json:"id"`
	Title     string    `validate:"required" json:"title"`
	Author    string    `validate:"required" json:"author"`
	Tags      string    `validate:"required" json:"tags"`
	Content   []uint8   `validate:"required" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Blogs) TableName() string {
	return "blogs"
}
