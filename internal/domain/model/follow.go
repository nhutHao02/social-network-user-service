package model

type FollowRequest struct {
	FollowerID  uint64 `json:"followerID" db:"FollowerID"`
	FollowingID uint64 `json:"followingID" db:"FollowingID"`
}

type FollowIDParam struct {
	ID uint64 `uri:"id"`
}

type FollowResponse struct {
	UserInfo []FollowUserInfoResponse `json:"followUserInfo"`
	Total    int                      `json:"total"`
}
