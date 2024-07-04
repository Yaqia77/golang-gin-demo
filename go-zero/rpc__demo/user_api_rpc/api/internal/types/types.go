// Code generated by goctl. DO NOT EDIT.
package types

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoRequest struct {
	ID uint `path:"id"`
}

type UserInfoResponse struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
}