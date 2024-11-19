package entity

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type User struct {
	ID            uint64      `json:"id" db:"ID"`
	Email         string      `json:"email" db:"Email"`
	Password      string      `json:"password" db:"Password"`
	FullName      null.String `json:"fullName" db:"FullName"`
	Sex           null.Bool   `json:"sex" db:"Sex"`
	Bio           null.String `json:"bio" db:"Bio"`
	UrlAvt        null.String `json:"urlAvt" db:"UrlAvt"`
	UrlBackground null.String `json:"urlBackground" db:"UrlBackground"`
	CreatedAt     time.Time   `json:"createdAt" db:"CreatedAt"`
	UpdatedAt     time.Time   `json:"updatedAt" db:"UpdatedAt"`
	DeletedAt     *time.Time  `json:"deletedAt" db:"DeletedAt"`
	LocationID    null.Int    `json:"locationID" db:"LocationID"`
}
