package models

type SocialMedia struct {
	ID             uint   `gorm:"primary_key" json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	User_ID        uint   `json:"user_id"`
}
