// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoResponse struct {
	SuerId   uint   `json:"userId"`
	Username string `json:"username"`
}
