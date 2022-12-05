package model

type ParamSignUp struct {
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
