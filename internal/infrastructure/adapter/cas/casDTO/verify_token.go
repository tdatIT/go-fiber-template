package casDTO

func (CASVerifyTokenReq) URL() string {
	return "/rest/v1/auth/verify"
}

type CASVerifyTokenReq struct {
	Token string `json:"token"`
}

type CASVerifyTokenResp struct {
	UserId   int      `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}
