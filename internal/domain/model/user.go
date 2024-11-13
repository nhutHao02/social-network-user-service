package model

import (
	"time"
)

type UserParam struct {
	ID int `uri:"id"`
}

type UserInfoResponse struct {
	ID               int               `json:"id" db:"ID"`
	Email            string            `json:"email" db:"Email"`
	FullName         *string           `json:"fullName" db:"FullName"`
	Sex              *bool             `json:"sex" db:"Sex"`
	Bio              *string           `json:"bio" db:"Bio"`
	UrlAvt           *string           `json:"urlAvt" db:"UrlAvt"`
	UrlBackground    *string           `json:"urlBackground" db:"UrlBackground"`
	CreatedAt        time.Time         `json:"createdAt" db:"CreatedAt"`
	UpdatedAt        time.Time         `json:"updatedAt" db:"UpdatedAt"`
	DeletedAt        *time.Time        `json:"deletedAt" db:"DeletedAt"`
	LocationResponse *LocationResponse `json:"location" db:"location"`
}

type UserUpdateRequest struct {
	ID            int                    `json:"id" db:"ID"`
	FullName      *string                `json:"fullName" db:"FullName"`
	Sex           *bool                  `json:"sex" db:"Sex"`
	Bio           *string                `json:"bio" db:"Bio"`
	UrlAvt        *string                `json:"urlAvt" db:"UrlAvt"`
	UrlBackground *string                `json:"urlBackground" db:"UrlBackground"`
	Location      *LocationUpdateRequest `json:"location"`
}

type UserUpdatePassRequest struct {
	ID       int    `json:"id" db:"ID"`
	Password string `json:"password" db:"Password"`
}

type FollowUserInfoResponse struct {
	ID            int        `json:"id" db:"ID"`
	Email         string     `json:"email" db:"Email"`
	FullName      *string    `json:"fullName" db:"FullName"`
	UrlAvt        *string    `json:"urlAvt" db:"UrlAvt"`
	UrlBackground *string    `json:"urlBackground" db:"UrlBackground"`
	CreatedAt     time.Time  `json:"createdAt" db:"CreatedAt"`
	UpdatedAt     time.Time  `json:"updatedAt" db:"UpdatedAt"`
	DeletedAt     *time.Time `json:"deletedAt" db:"DeletedAt"`
}
