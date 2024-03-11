package token

import "time"

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *CustomClaim, error)
	VerifyToken(token string) (*CustomClaim, error)
}
