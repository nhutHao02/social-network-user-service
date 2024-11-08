package entity

import (
	"time"
)

type Location struct {
	ID          int        `json:"id" db:"ID"`
	City        string     `json:"city" db:"City"`
	District    string     `json:"district" db:"District"`
	Ward        string     `json:"ward" db:"Ward"`
	Description string     `json:"description" db:"Description"`
	CreatedAt   time.Time  `json:"createdAt" db:"CreatedAt"`
	UpdatedAt   time.Time  `json:"updatedAt" db:"UpdatedAt"`
	DeletedAt   *time.Time `json:"deletedAt" db:"DeletedAt"`
}
