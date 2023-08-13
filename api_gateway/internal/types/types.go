// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `form:"username"` // 注册用户名，最长32个字符
	Password string `form:"password"` // 密码，最长32个字符
}

type LoginResp struct {
	StatusCode int64  `json:"status_code"`         // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,optional"` // 返回状态描述
	Token      string `json:"token"`               // 用户鉴权token
	UserID     int64  `json:"user_id"`             // 用户id
}

type CommentActionReq struct {
	ActionType  int32  `form:"action_type"`           // 1-发布评论，2-删除评论
	CommentID   int64  `form:"comment_id,optional"`   // 要删除的评论id，在action_type=2的时候使用
	CommentText string `form:"comment_text,optional"` // 用户填写的评论内容，在action_type=1的时候使用
	Token       string `form:"token"`                 // 用户鉴权token
	VideoID     int64  `form:"video_id"`              // 视频id
}

type CommentActionResp struct {
	Comment    Comment `json:"comment"`              // 评论成功返回评论内容，不需要重新拉取整个列表
	StatusCode int64   `json:"status_code,optional"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`           // 返回状态描述
}

type Comment struct {
	ID         int64  `form:"id"`          // 评论id
	Content    string `form:"content"`     // 评论内容
	User       User   `form:"user"`        // 评论用户信息
	CreateDate string `form:"create_date"` // 评论发布日期，格式 mm-dd
}

type User struct {
	ID             int64  `json:"id"`                       // 用户id
	Name           string `json:"name"`                     // 用户名称
	Signature      string `json:"signature,optional"`       // 个人简介
	TotalFavorited string `json:"total_favorited,optional"` // 获赞数量
	WorkCount      int64  `json:"work_count,optional"`      // 作品数
}
