package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

type JWTData struct {
	Username string
	IssuedAt time.Time
	ExpirationTime time.Time
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(username string) (string, error) {
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": now.Unix(),
		"exp": now.Add(time.Minute * 10).Unix(),
	})

	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *JWT) Parse(token string) (*JWTData, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims := t.Claims.(jwt.MapClaims)

	sub, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}

	iat, err := claims.GetIssuedAt()
	if err != nil {
		return nil, err
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	return &JWTData{
		Username: sub,
		IssuedAt: iat.Time,
		ExpirationTime: exp.Time,
	}, nil
}
