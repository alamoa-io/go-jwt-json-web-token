package auth

import (
	"crypto/rsa"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"

	"go-jwt-json-web-token/models"
)

const (
	Subject                = "AccessToken"
	Issuer                 = "https://idp.alamoa.io"
	Audience               = "https://api.alamoa.io"
	TokenExpiration        = time.Minute * time.Duration(1)
	RefreshTokenExpiration = time.Hour * time.Duration(1)
	tokenPrivateKey        = "./auth/keys/token.rsa"             // openssl genrsa -out token.rsa
	TokenPublicKey         = "./auth/keys/token.rsa.pub"         // openssl rsa -in token.rsa -pubout > token.rsa.pub
	refreshTokenPrivateKey = "./auth/keys/refresh_token.rsa"     // openssl genrsa -out refresh_token.rsa
	RefreshTokenPublicKey  = "./auth/keys/refresh_token.rsa.pub" // openssl rsa -in refresh_token.rsa -pubout > refresh_token.rsa.pub
)

var (
	tokenSignKey *rsa.PrivateKey
	// If you want to sign with a string instead of RSA
	// tokenSignKey = []byte("secretkey")
	TokenVerifyKey        *rsa.PublicKey
	refreshTokenSignKey   *rsa.PrivateKey
	RefreshTokenVerifyKey *rsa.PublicKey
)

func init() {
	signBytes, _ := os.ReadFile(tokenPrivateKey)
	tokenSignKey, _ = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	verifyBytes, _ := os.ReadFile(TokenPublicKey)
	TokenVerifyKey, _ = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	refreshSignBytes, _ := os.ReadFile(refreshTokenPrivateKey)
	refreshTokenSignKey, _ = jwt.ParseRSAPrivateKeyFromPEM(refreshSignBytes)
	refreshTokenVerifyBytes, _ := os.ReadFile(RefreshTokenPublicKey)
	RefreshTokenVerifyKey, _ = jwt.ParseRSAPublicKeyFromPEM(refreshTokenVerifyBytes)
}

type Claim struct {
	Email string
	jwt.StandardClaims
}

func NewClaim(email string) *Claim {
	return &Claim{
		Email: email,
	}
}

func (c *Claim) GenerateToken() (token string, err error) {
	claims := jwt.StandardClaims{
		Issuer:    Issuer,
		Subject:   Subject,
		Audience:  Audience,
		IssuedAt:  time.Now().Local().Unix(),
		ExpiresAt: time.Now().Local().Add(TokenExpiration).Unix(),
	}
	if err = copier.CopyWithOption(c, claims, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return "", err
	}

	// If you want to sign with a string with HS256 instead of RSA
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	generatedToken := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	token, err = generatedToken.SignedString(tokenSignKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *Claim) GenerateRefreshToken() (token string, err error) {
	claims := jwt.StandardClaims{
		Id:        uuid.New().String(),
		Issuer:    Issuer,
		Subject:   Subject,
		Audience:  Audience,
		IssuedAt:  time.Now().Local().Unix(),
		ExpiresAt: time.Now().Local().Add(RefreshTokenExpiration).Unix(),
	}
	if err = copier.CopyWithOption(c, claims, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return "", err
	}

	// If you want to sign with a string with HS256 instead of RSA
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	token, err = refreshToken.SignedString(refreshTokenSignKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *Claim) UpdateRefreshToken() (token string, err error) {
	claims := jwt.StandardClaims{
		Issuer:   Issuer,
		Subject:  Subject,
		Audience: Audience,
	}
	if err = copier.CopyWithOption(c, claims, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return "", err
	}
	c.Id = uuid.New().String()
	c.IssuedAt = time.Now().Local().Unix()
	// If you want to log in again after a certain period of time from when you first logged in, use the old refresh
	// token time. If you want to newly extend expire, set it as follows.
	// claim.ExpiresAt = time.Now().Local().Unix()

	// If you want to sign with a string instead of RSA
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, updateClaims)
	updatedToken := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	token, err = updatedToken.SignedString(refreshTokenSignKey)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (c *Claim) SetBlackListToken() {
	now := time.Now()
	tokenExpiresAt := time.Unix(c.ExpiresAt, 0).Sub(now)
	// You can also check if there are any users who repeatedly log in and log out during other processes
	// by entering their email here.
	response, err := models.REDIS.Set(c.Id, c.Email, tokenExpiresAt).Result()
	log.Println(response, err)
}

func (c *Claim) IsBlackListToken() bool {
	expiresAt, err := models.REDIS.Get(c.Id).Result()
	if err != nil {
		log.Println(err)
	}
	if expiresAt != "" {
		return true
	}
	return false
}

func ValidateToken(token string, verifyKey *rsa.PublicKey) (*Claim, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*Claim)
	if !ok {
		err = errors.New("could not parse claims")
		return nil, err
	}

	return claims, nil
}
