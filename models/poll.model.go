package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Poll struct {
	ID          uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	User        uuid.UUID   `gorm:"not null" json:"user,omitempty"`
	Description string      `gorm:"not null" json:"description,omitempty"`
	Options     StringArray `gorm:"type:text[];not null" json:"options,omitempty"`
	Counts      StringArray `gorm:"type:text[];not null" json:"counts,omitempty"`
	CreatedAt   time.Time   `gorm:"not null" json:"created_at,omitempty"`
}

type PollRequest struct {
	Description string   `gorm:"not null" json:"description,required"`
	Options     []string `gorm:"not null" json:"options,required"`
}

type StringArray []string

func (a StringArray) Value() (driver.Value, error) {
	// Convert the slice to a PostgreSQL array string
	if len(a) == 0 {
		return "{}", nil
	}
	return "{" + strings.Join(a, ",") + "}", nil
}

func (a *StringArray) Scan(value interface{}) error {
	// Convert the PostgreSQL array string to a slice
	strVal, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan StringArray")
	}

	*a = strings.Split(strings.Trim(strVal, "{}"), ",")
	return nil
}
