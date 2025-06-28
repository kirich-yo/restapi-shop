package jwt

import (
	"time"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

type JWTData struct {
	UserID uint
	IssuedAt time.Time
	ExpirationTime time.Time
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(userID uint, lifetime time.Duration) (string, error) {
	now := time.Now()

	sub := strconv.FormatUint(uint64(userID), 10)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": now.Unix(),
		"exp": now.Add(lifetime).Unix(),
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

	userID, err := strconv.ParseUint(sub, 10, 32)

	return &JWTData{
		UserID: uint(userID),
		IssuedAt: iat.Time,
		ExpirationTime: exp.Time,
	}, nil
}
