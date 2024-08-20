package oauth2

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// SocialIDs 用于存储第三方登录 ID
type OAuthAccounts struct {
	GithubID *string `gorm:"column:github_id;unique" json:"github_id,omitempty"` // GitHub 用户 ID
	GoogleID *string `gorm:"column:google_id;unique" json:"google_id,omitempty"` // Google 用户 ID
	WechatID *string `gorm:"column:wechat_id;unique" json:"wechat_id,omitempty"` // 微信用户 ID
	QQID     *string `gorm:"column:qq_id;unique" json:"qq_id,omitempty"`         // QQ 用户 ID
}

var googleOauthConfig = oauth2.Config{
	ClientID:     "YOUR_GOOGLE_CLIENT_ID",
	ClientSecret: "YOUR_GOOGLE_CLIENT_SECRET",
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}
