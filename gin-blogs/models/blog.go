package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	ID         int    `json:"id" gorm:"primarykey;autoIncrement" `
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	AuthorID   int    `json:"author_id" binding:"required"`
	CreatedAt  string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  string `json:"updated_at" gorm:"autoCreateTime"`
	CoverImage string `json:"cover_image"`
}

func AutoMigrateBlogTable(db *gorm.DB) error {
	if err := db.AutoMigrate(&Blog{}); err != nil {
		return err
	}
	return nil
}
