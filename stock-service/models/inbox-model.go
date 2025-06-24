package models

import "time"

type InboxMessage struct {
	ID          uint   `gorm:"primaryKey"`
	MessageID   string `gorm:"uniqueIndex"`
	Payload     []byte
	Processed   bool
	ProcessedAt *time.Time
	CreatedAt   time.Time
}
