package schemas

type RequestAuthForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type ResponseToken struct {
	Token string `json:"token" binding:"required"`
	Type  string `json:"type"`
}

func NewBearerToken(token string) *ResponseToken {
	return &ResponseToken{
		Token: token,
		Type:  "Bearer",
	}
}
