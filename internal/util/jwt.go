package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
	Email     string    `json:"email"`
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type JWTMaker struct {
	secretKey        string
	refreshSecretKey string
}

func (p Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewJWTMaker(secretKey string, refreshSecretKey string) *JWTMaker {
	return &JWTMaker{
		secretKey:        secretKey,
		refreshSecretKey: refreshSecretKey,
	}
}

func (m *JWTMaker) CreateTokenPair(email string, duration time.Duration) (TokenPair, error) {
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	accessTokenPayload := Payload{
		Email:     email,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}

	refreshTokenPayload := Payload{
		Email:     email,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt.Add(time.Hour * 24 * 7),
	}

	accessToken, err := m.createToken(accessTokenPayload, m.secretKey)
	if err != nil {
		return TokenPair{}, err
	}

	refreshToken, err := m.createToken(refreshTokenPayload, m.refreshSecretKey)
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (m *JWTMaker) createToken(payload Payload, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(secretKey))
}

func (m *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
