// Code generated by goctl. DO NOT EDIT.
package types

type RegisterOrLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserReq struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type DataReply struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type UserReply struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	User       User   `json:"user"`
}

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}