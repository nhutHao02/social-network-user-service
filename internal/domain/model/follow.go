package model

type FollowRequest struct {
	FollowerID  int `json:"followerID" db:"FollowerID"`
	FollowingID int `json:"followingID" db:"FollowingID"`
}
