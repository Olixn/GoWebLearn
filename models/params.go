package models

// 定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVote 投票请求参数
type ParamVote struct {
	UserID    int64  `json:"user_id,string"`                          // 用户ID
	PostID    string `json:"post_id" binding:"required"`              // 帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票（1）反对票（-1）取消投票（0）
}

type ParamPostList struct {
	Page        int64  `form:"page"`
	Size        int64  `form:"size"`
	Order       string `form:"order"`
	CommunityID int64  `json:"community_id" form:"community_id"'`
}

type ParamCommunityPostList struct {
	*ParamPostList
}
