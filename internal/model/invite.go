package model

// Invite представление приглашения на встречу
type Invite struct {
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:user_id"`
	MeetingID uint    `gorm:"not null"`
	Meetings  Meeting `gorm:"foreignKey:meeting_id"`
	Approved  bool    `gorm:"not null"`
}
