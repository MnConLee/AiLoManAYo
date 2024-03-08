package form

type User struct {
	InviteId   string `form:"Invite_id" `    // 邀请的人id
	MagicBall  int64  `form:"magic_ball"`    // 魔法球
	NickName   string `form:"nick_name" `    // 微信昵称
	UserName   string `form:"user_name" `    // 微信用户唯一标识
	City       string `form:"city" `         // 用户所在城市
	HeadImgUrl string `json:"head_img_url" ` // 用户头像
	Sex        int    `json:"sex" `          // 性别 0：男 1：女
}
