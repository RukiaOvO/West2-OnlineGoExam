package types

type RegisterRequest struct {
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"pass_word" form:"pass_word"`
}

type LoginRequest struct {
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"pass_word" form:"pass_word"`
}

type UserLoginResponse struct {
	Id        int64  `json:"id"`
	UserName  string `json:"user_name"`
	CreatedAt int64  `json:"created_at"`
}

type TokenData struct {
	User        interface{} `json:"user"`
	AccessToken string      `json:"access_token"`
}
