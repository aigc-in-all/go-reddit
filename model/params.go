package model

type ParamSignUp struct {
	UserName   string `json:"user_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`              // 帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1)，反对票(-1)
}
