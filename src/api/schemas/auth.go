package schemas

type RequestAuthForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type BearerToken struct {
	Token string `json:"token" binding:"required"`
	Type  string `json:"type"`
}

func NewBearerToken(token string) *BearerToken {
	return &BearerToken{
		Token: token,
		Type:  "Bearer",
	}
}
