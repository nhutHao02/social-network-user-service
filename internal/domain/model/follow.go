package model

type FollowRequest struct {
	FollowerID  int `json:"followerID" db:"FollowerID"`
	FollowingID int `json:"followingID" db:"FollowingID"`
}

type FollowIDParam struct {
	ID int `uri:"id"`
}

type FollowResponse struct {
	UserInfo []FollowUserInfoResponse `json:"followUserInfo"`
	Total    int                      `json:"total"`
}
