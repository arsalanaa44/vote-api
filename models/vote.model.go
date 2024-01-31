package models

import (
	"time"

	"github.com/google/uuid"
)

type Vote struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	User      uuid.UUID `gorm:"not null" json:"user,omitempty"`
	Poll      uuid.UUID `gorm:"not null" json:"poll,omitempty"`
	Choice    int       `gorm:"not null" validate:"int" json:"choice,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
}

type VoteRequest struct {
	Poll   uuid.UUID `gorm:"not null" json:"poll,required"`
	Choice int       `gorm:"not null" validate:"int" json:"choice,required"`
}
