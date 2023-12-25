package consts

import "time"

const (
	JwtSecret              = "IamAbakaya"
	AccessTokenExpireTime  = time.Hour * 3
	RefreshTokenExpireTime = time.Hour * 3
	AccessIssuer           = "baka_accesstoken"
	RefreshIssuer          = "baka_refreshtoken"
	PasswordCost           = 12
)
